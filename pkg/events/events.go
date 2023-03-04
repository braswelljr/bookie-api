package events

import "encore.dev/pubsub"

// DeleteAllUserTasks - Event to delete all user tasks
var DeleteAllUserTasks = pubsub.NewTopic[*DeleteAllUserTasksEvent]("delete-all-user-tasks", pubsub.TopicConfig{
	DeliveryGuarantee: pubsub.AtLeastOnce,
})
