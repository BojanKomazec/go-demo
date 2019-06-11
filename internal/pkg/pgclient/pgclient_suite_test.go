package pgclient

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BojanKomazec/go-demo/internal/pkg/dbclient"
)

func TestPgclient(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "PgClient Suite")
}

var _ = Describe("New", func() {
	dbConnParams := dbclient.NewConnParams(
		uuid.New().String(),
		rand.Intn(65535),
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String())

	Describe("Returns PgClient", func() {
		pgClient, _ := newImpl(dbConnParams)

		It("should have Correct DriverName", func() {
			Expect(pgClient.driverName).To(Equal("postgres"))
		})

		It("should have Correct DataSourceName", func() {
			Expect(pgClient.dataSourceName).To(
				Equal(fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
					dbConnParams.Host(), dbConnParams.Port(), dbConnParams.User(),
					dbConnParams.Password(), dbConnParams.DbName())))
		})
	})
})
