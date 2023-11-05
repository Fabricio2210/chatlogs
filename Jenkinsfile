pipeline {
    agent any
    stages {
        stage('Checkout Codebase') {
            steps {
                checkout scm: [$class: 'GitSCM', branches: [[name: '*/main']], userRemoteConfigs: [[credentialsId: 'github-ssh-key', url: 'git@github.com:Fabricio2210/chatlogs.git']]
            }
        }
        stage('Set Go Environment') {
            steps {
                script {
                    def goPath = "/var/lib/jenkins/workspace/chatLogs/go"  // Customize the path
                    withEnv(["GOPATH=${goPath}"]) {
                        sh "mkdir -p ${goPath}/src ${goPath}/bin ${goPath}/pkg"
                    }
                }
            }
        }
        stage('Build') {
            steps {
                echo 'Building Codebase'
                script {
                    sh 'ls -l' // List files in the workspace
                    sh 'go build -o myprogram main.go'
                    sh 'ls -l' // List files in the workspace again to check if the executable is created
                }
            }
        }
        stage('Deploy') {
            steps {
                echo 'Executing the program'
                sh './myprogram'
            }
        }
    }
    post {
        always {
            echo 'Done'
        }
    }
}

