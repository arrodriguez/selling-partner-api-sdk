module github.com/arrodriguez/selling-partner-api-sdk

go 1.16

require (
	github.com/aws/aws-sdk-go v1.36.23
	github.com/dpoetzschke/selling-partner-api-sdk v0.0.0-20230306125519-9be6dcb3e738
	github.com/google/uuid v1.1.4
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.6.1
)

replace github.com/dpoetzschke/selling-partner-api-sdk v0.0.0-20230306125519-9be6dcb3e738 => github.com/arrodriguez/selling-partner-api-sdk v0.0.0-20230318200117-ae816db60360
