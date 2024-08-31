# Recommendation Engine

This project is a microservice-based application designed to track user behavior and provide personalized recommendations based on that behavior. The application leverages a combination of Golang, Flask, and Kafka to deliver a scalable and efficient recommendation system.



## Architecture

- User-Tracking Microservice: This is used to track every single action made by the user and stream it to kafka
- Recommendation Microservice: This microservice consumes the events from kafka, gets a prediction from the flask microservice and stores it in the DB
- Flask Microservice: This microservice acts as the entry point to the machine learning modela and returns a prediction 


## Features

- User Behavior Tracking: Captures and streams user interactions to Kafka.
- Real-Time Recommendations: Processes user behavior data and generates personalized recommendations.
- Microservices Architecture: Each component is modular and independently deployable, allowing for easy scaling and maintenance.
