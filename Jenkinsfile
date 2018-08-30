#!/usr/bin/env groovy

pipeline {
    environment {
        DOCKER_HUB_USER = 'ekinyucel'
        DOCKER_CREDENTIAL_ID = 'docker_hub'
    }

    agent none

    stages {
        stage('Stop running container') {
            agent any
            sh "docker rm $(docker ps -a -q --filter ancestor=rest_api)"
            echo "Running application forced to stopped and pruned"
        }
        stage('Build') {
            agent any
            steps {
                git "https://github.com/Broke116/go_rest_api.git"
                sh 'docker build -t rest_api .'
            }
        }
        stage('Docker Run'){
            agent any
            steps {
                sh "docker run -d --rm -p 4000:3030 -t rest_api"
                echo "Application started on port: 4000"
            }
        }
    }
}