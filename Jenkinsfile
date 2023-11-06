pipeline {
    agent any

    stages {
         stage('Checkout Codebase'){
            steps{
                checkout scm: [$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[credentialsId: 'github-ssh-key', url: 'git@github.com:Fabricio2210/chatlogs.git']]] 
            }
        }

        stage('Build') {
            steps {
                sh 'go get -v' // Download dependencies
                sh 'go build -o chatlogs' // Build your Go application
            }
        }

        stage('Deploy') {
            steps {
                // You can deploy the Go application to your server here
                // For simplicity, we'll just print a message
                echo 'Deploying the application...'
            }
        }
    }
}



