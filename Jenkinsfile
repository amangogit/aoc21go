pipeline{
    // agent any
    agent {
        docker { image 'golang:1.14' }
    }
    stages{
        stage('Test'){
            steps{
                sh 'echo $PATH'
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