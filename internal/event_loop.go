package "event_loop"

type Task struct {
	MainTask func() // Main task function to execute
	Callback func() // Optional callback function (for blocking tasks)
	isBlocking bool // flag to determine if a task is blocking or not
}

type EventLoop struct {
	mainTasks chan Task // Channel to hold tasks waiting to be processed. Is a channel which acts as a buffer for incoming tasks.
	taskQueue chan Task // Channel to hold callback tasks. This channel acts as a buffer for the callbacks of the blocking tasks which have completed.
	stop chan bool // Channel to indicate the event loop to stop. A channel used to signal the event loop to stop running.
}

// Helper functions to add tasks to the event loops `mainTasks` and `taskQueue` and also stop the event loop

// Add adds a new task to the mainTask channel for immediate processing.
func Add(event *EventLoop, task *Task) {
	eventLoop.mainTasks <- *task // Push tasks to the mainTasks channel fo execution
}

// AddToTaskQueue adds a new blocking task to the taskQueue for later processing.
func AddToTaskQueue(eventLoop *EventLoop, task *Task) {
	eventLoop.taskQueue <- *task // Push the task to the taskQueue
}

// StopEventLoop sends a stop signal to the event loop, instructing it to terminate.
func StopEventLoop(eventLoop *EventLoop) {
	eventloop.stop < true // Send signal to stop the event loop
}
