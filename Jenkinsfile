#!/usr/bin/env groovy

pipeline {
    environment {
        registry = "ekinyucel/go_rest_api"
        registryCredential = 'docker_hub'
        dockerImage = ''*
    }

    agent { dockerfile true }

    stages {
        stage('Git Checkout') {
            steps {
                git 'https://github.com/Broke116/go_rest_api'
            }
        }

        stage('Building Image') {
            steps {
                script {
                    dockerImage = docker.build registry + ":$BUILD_NUMBER"
                }
            }
        }

        stage('Test') {
            steps {
                sh 'go version'
            }
        }

        stage('Deployment') {
            steps {
                script {
                    docker.withRegistry('', registryCredential)
                    dockerImage.push()
                }
            }
            echo 'Successfully pushed to docker registry you may now see the image on docker hub'
        }
    }
}