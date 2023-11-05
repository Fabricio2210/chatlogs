pipeline {
    agent any
    stages {
        stage('Checkout Codebase') {
            steps {
                checkout scm: [$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[credentialsId: 'github-ssh-key', url: 'git@github.com:Fabricio2210/chatlogs.git']]
            }
        }
        stage('Build') {
            steps {
                echo 'Building Codebase'
                sh 'go build -o main main.go'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Executing the program'
                sh './main'
            }
        }
    }
    post {
        always {
            echo 'Done'
        }
    }
}
