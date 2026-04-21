package queue

//	!IMP : this is just a placeholder
//
// TODO : fill struct with actual data required for the job
type PullJobData struct {
	Token  string
	Owner  string
	Repo   string
	Branch string
}

type BuildJobData struct {
	Token  string
	Owner  string
	Repo   string
	Branch string
}

type DeployJobData struct {
	Token  string
	Owner  string
	Repo   string
	Branch string
}

type JobQueue struct {
	PullQueue   chan *PullJobData
	BuildQueue  chan *BuildJobData
	DeployQueue chan *DeployJobData
}

// initializes the job queues
func InitWorkerQueue() *JobQueue {
	pull := make(chan *PullJobData, 10)
	build := make(chan *BuildJobData, 10)
	deploy := make(chan *DeployJobData, 10)

	return &JobQueue{
		PullQueue:   pull,
		BuildQueue:  build,
		DeployQueue: deploy,
	}
}

// push job to pull worker queue
func (j *JobQueue) EnqueuePullJob(data *PullJobData) {
	j.PullQueue <- data
}

// push job to build worker queue
func (j *JobQueue) EnqueueBuildJob(data *BuildJobData) {
	j.BuildQueue <- data
}

// push job to deploy worker queue
func (j *JobQueue) EnqueueDeployJob(data *DeployJobData) {
	j.DeployQueue <- data
}

// closes all the queue channels
func (j *JobQueue) CloseQueue() {
	close(j.PullQueue)
	close(j.BuildQueue)
	close(j.DeployQueue)
}
