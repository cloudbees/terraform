// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudformation provides the client and types for making API
// requests to AWS CloudFormation.
//
// AWS CloudFormation allows you to create and manage AWS infrastructure deployments
// predictably and repeatedly. You can use AWS CloudFormation to leverage AWS
// products, such as Amazon Elastic Compute Cloud, Amazon Elastic Block Store,
// Amazon Simple Notification Service, Elastic Load Balancing, and Auto Scaling
// to build highly-reliable, highly scalable, cost-effective applications without
// creating or configuring the underlying AWS infrastructure.
//
// With AWS CloudFormation, you declare all of your resources and dependencies
// in a template file. The template defines a collection of resources as a single
// unit called a stack. AWS CloudFormation creates and deletes all member resources
// of the stack together and manages all dependencies between the resources
// for you.
//
// For more information about AWS CloudFormation, see the AWS CloudFormation
// Product Page (http://aws.amazon.com/cloudformation/).
//
// Amazon CloudFormation makes use of other AWS products. If you need additional
// technical information about a specific AWS product, you can find the product's
// technical documentation at docs.aws.amazon.com (http://docs.aws.amazon.com/).
//
// See https://docs.aws.amazon.com/goto/WebAPI/cloudformation-2010-05-15 for more information on this service.
//
// See cloudformation package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/
//
// Using the Client
//
// To use the client for AWS CloudFormation you will first need
// to create a new instance of it.
//
// When creating a client for an AWS service you'll first need to have a Session
// already created. The Session provides configuration that can be shared
// between multiple service clients. Additional configuration can be applied to
// the Session and service's client when they are constructed. The aws package's
// Config type contains several fields such as Region for the AWS Region the
// client should make API requests too. The optional Config value can be provided
// as the variadic argument for Sessions and client creation.
//
// Once the service's client is created you can use it to make API requests the
// AWS service. These clients are safe to use concurrently.
//
//   // Create a session to share configuration, and load external configuration.
//   sess := session.Must(session.NewSession())
//
//   // Create the service's client with the session.
//   svc := cloudformation.New(sess)
//
// See the SDK's documentation for more information on how to use service clients.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws package's Config type for more information on configuration options.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the AWS CloudFormation client CloudFormation for more
// information on creating the service's client.
// https://docs.aws.amazon.com/sdk-for-go/api/service/cloudformation/#New
//
// Once the client is created you can make an API request to the service.
// Each API method takes a input parameter, and returns the service response
// and an error.
//
// The API method will document which error codes the service can be returned
// by the operation if the service models the API operation's errors. These
// errors will also be available as const strings prefixed with "ErrCode".
//
//   result, err := svc.CancelUpdateStack(params)
//   if err != nil {
//       // Cast err to awserr.Error to handle specific error codes.
//       aerr, ok := err.(awserr.Error)
//       if ok && aerr.Code() == <error code to check for> {
//           // Specific error code handling
//       }
//       return err
//   }
//
//   fmt.Println("CancelUpdateStack result:")
//   fmt.Println(result)
//
// Using the Client with Context
//
// The service's client also provides methods to make API requests with a Context
// value. This allows you to control the timeout, and cancellation of pending
// requests. These methods also take request Option as variadic parameter to apply
// additional configuration to the API request.
//
//   ctx := context.Background()
//
//   result, err := svc.CancelUpdateStackWithContext(ctx, params)
//
// See the request package documentation for more information on using Context pattern
// with the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/request/
package cloudformation
