package agent

import (
	"sync"
	"time"

	"github.com/Axway/agent-sdk/pkg/jobs"
	hc "github.com/Axway/agent-sdk/pkg/util/healthcheck"
)

// constants for retry interval for stream job
const (
	defaultRetryInterval   = 5 * time.Second
	maxRetryInterval       = 5 * time.Minute
	maxNumRetryForInterval = 3
)

// eventsJob interface for a job to execute to retrieve events in either stream or poll mode
type eventsJob interface {
	Start() error
	Status() error
	Stop()
	Healthcheck(_ string) *hc.Status
}

// eventProcessorJob job wrapper for a streamerClient that starts a stream and an event manager.
type eventProcessorJob struct {
	streamer      eventsJob
	stop          chan interface{}
	jobID         string
	retryInterval time.Duration
	numRetry      int
	name          string
	mutex         sync.RWMutex
}

// newEventProcessorJob creates a job for the streamerClient
func newEventProcessorJob(eventJob eventsJob, name string) jobs.Job {
	streamJob := &eventProcessorJob{
		streamer:      eventJob,
		stop:          make(chan interface{}, 1),
		retryInterval: defaultRetryInterval,
		numRetry:      0,
		name:          name,
	}

	jobID, _ := jobs.RegisterDetachedChannelJobWithName(streamJob, streamJob.stop, name)
	streamJob.mutex.Lock()
	streamJob.jobID = jobID
	streamJob.mutex.Unlock()

	jobs.RegisterIntervalJobWithName(newCentralHealthCheckJob(eventJob), time.Second*3, "Central Health Check")

	return streamJob
}

// Execute starts the stream
func (j *eventProcessorJob) Execute() error {
	go func() {
		<-j.stop
		j.streamer.Stop()
		j.renewRegistration()
	}()

	return j.streamer.Start()
}

// Status gets the status
func (j *eventProcessorJob) Status() error {
	status := j.streamer.Status()
	if status == nil {
		j.retryInterval = defaultRetryInterval
		j.numRetry = 0
	}
	return status
}

// Ready checks if the job to start the stream is ready
func (j *eventProcessorJob) Ready() bool {
	return true
}

func (j *eventProcessorJob) renewRegistration() {
	j.mutex.RLock()
	jobID := j.jobID
	j.mutex.RUnlock()

	defer time.AfterFunc(j.retryInterval, func() {
		jobID, _ := jobs.RegisterDetachedChannelJobWithName(j, j.stop, j.name)
		j.mutex.Lock()
		defer j.mutex.Unlock()
		j.jobID = jobID
	})

	if jobID != "" {
		j.mutex.Lock()
		defer j.mutex.Unlock()

		jobs.UnregisterJob(j.jobID)
		j.jobID = ""
		j.numRetry++
		if j.numRetry == maxNumRetryForInterval {
			j.numRetry = 0
			j.retryInterval = j.retryInterval * 2
			if j.retryInterval > maxRetryInterval {
				j.retryInterval = defaultRetryInterval
			}
		}
	}
}
