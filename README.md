# Go Learn Go

This repo contains a curriculum for a book club on Go at Code Climate. It may
change as we find new things we want to know.

1. Installing Go

   - Work through [Installation]
   - Read through the [workspace] explanation

1. Go basics

   - Work through [Learning Go]

1. Testing

   - Read through [Testing in Go]
   - Write tests for something you did in a previous section

1. Web apps

   - Read through [Creating a RESTful API With Golang]
   - Build an API that returns the weather as JSON when responding to
     `localhost:8081/10011` where the path is the zip code
   - Deploy that web app to Heroku

1. Kafka

   - Read up on [confluent-kafka-go], including the examples linked from the
     README
   - Build 2 programs. One should produce a never-ending sequence of numbers to
     a Kafka topic. The other should consume that topic and print out the
     messages it is consuming.
   - Update the programs to work with a topic of multiple partitions.

1. Docker

   - Package the programs from the Kafka section as Docker containers

1. Monitoring

   - Update your weather API to report metrics to Datadog. How long does each
     request take? What portion of that request is spent getting the weather
     from an external API?

[confluent-kafka-go]: https://github.com/confluentinc/confluent-kafka-go
[Creating a RESTful API with Golang]: https://tutorialedge.net/golang/creating-restful-api-with-golang/
[Installation]: https://astaxie.gitbooks.io/build-web-application-with-golang/en/01.1.html
[Learning Go]: https://www.miek.nl/go
[Testing in Go]: https://blog.codeship.com/testing-in-go/
[workspace]: https://astaxie.gitbooks.io/build-web-application-with-golang/en/01.2.html
