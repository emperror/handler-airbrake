// nolint: goconst
package airbrake_test

import (
	"github.com/airbrake/gobrake"

	"handler.emperror.dev/airbrake"
)

func ExampleNew() {
	projectID := int64(1)
	projectKey := "key"

	handler := airbrake.New(projectID, projectKey)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewFromNotifier() {
	projectID := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	handler := airbrake.NewFromNotifier(notifier)
	defer handler.Close() // Make sure to close the handler to flush all error reporting in progress

	// Output:
}

func ExampleNewSync() {
	projectID := int64(1)
	projectKey := "key"

	handler := airbrake.NewSync(projectID, projectKey)
	defer handler.Close()

	// Output:
}

func ExampleNewSyncFromNotifier() {
	projectID := int64(1)
	projectKey := "key"

	notifier := gobrake.NewNotifier(projectID, projectKey)
	handler := airbrake.NewSyncFromNotifier(notifier)
	defer handler.Close()

	// Output:
}
