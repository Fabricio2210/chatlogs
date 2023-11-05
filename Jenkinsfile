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
                script {
                    def buildOutput = sh(script: 'go build -o myprogram main.go', returnStatus: true)
                    if (buildOutput != 0) {
                        error "Build failed with exit code ${buildOutput}"
                    }
                }
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
