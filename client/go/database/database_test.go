package database

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/go-test/deep"
	"github.com/karagog/clock-go/simulated"
	"github.com/karagog/db-provider/server/lessor"
	"github.com/karagog/db-provider/server/lessor/databaseprovider/fake"
	pb "github.com/karagog/db-provider/server/proto"
	"github.com/karagog/db-provider/server/service"
	"github.com/karagog/db-provider/server/service/runner"
)

// Test that the client object successfully gets an instance from the service.
func TestDatabase(t *testing.T) {
	// Initialize a fake in-memory service.
	provider := &fake.DatabaseProvider{
		Info: pb.ConnectionInfo{
			AppConn: &pb.ConnectionDetails{
				User: "George",
			},
			RootConn: &pb.ConnectionDetails{
				User: "root",
			},
		},
	}
	svc := service.New(simulated.NewClock(time.Now()))
	l := lessor.New(provider, 1)
	svc.SetLessor(l)
	ctx := context.Background()
	go l.Run(ctx)

	r, err := runner.New(svc, "localhost:0")
	if err != nil {
		t.Fatal(err)
	}
	go r.Run()
	defer r.Stop()

	// Override the address so it uses our fake service.
	os.Setenv("DB_INSTANCE_PROVIDER_ADDRESS", r.Address())

	// Get a database instance.
	i := NewFromEnv(ctx)

	// Make sure we got the connection info given to us by the server.
	if diff := deep.Equal(i.Info, &provider.Info); diff != nil {
		t.Fatalf("Got unepected info: %v", diff)
	}

	i.Close() // nominal close

	// Close a second time, it should do nothing.
	i.Close()
}
