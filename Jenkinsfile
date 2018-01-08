node {
  stage('Build') {
    echo 'Building..'
    sh 'make all'
  }
  stage('Test') {
    echo 'Testing..'
    // sh 'make testverbose'
    sh 'coverageall'
  }
  stage('Deploy') {
    echo 'Deploying....'
    // NOP
  }
}