pipeline{
    agent any
    // agent {
    //     docker { image 'golang:1.14' }
    // }
    environment{
        USER = 'aman'
    }
    stages{
        stage('build'){
            steps{
                sh 'echo $PATH'
                sh 'pwd'
                sh 'echo $USER'
                sh '''
                    go version
                    ls -lah
                '''
                // retry(3){
                //     sh 'whoami'
                // }
                timeout(time: 3, unit: 'MINUTES'){
                    sh 'echo $USER'
                }
            }
        }
        stage('Test'){
            steps{
                sh 'go version'
                sh 'echo $USERNAME'
                echo "User env is ${USER}"
                sh 'getent passwd ${USER}'
                sh 'printenv'
            }
        }
    }
    post{
        always{
            echo 'This will always run'
        }
        success{
            echo 'This will run only if successful'
        }
        failure{
            echo 'This will run only if failed'
        }
        unstable{
            echo 'This will run only if the run was marked unstable'
        }
        changed{
            echo 'This will run only if the state of the Pipeline has changed'
            echo 'For example, if the Pipeline was previously failing but is now successful'
        }
    }
}