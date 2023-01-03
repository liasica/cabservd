node {
    stage('kxservd') {
        if (TAG == 'kxservd') {
            // timeout (time: 1, unit: 'HOURS' )  {
            //     input 'Deploy [kxservd]?'
            // }
            echo '开始部署[kxservd]'
            sshagent (credentials: ['Jenkins']) {
                sh "ssh -o StrictHostKeyChecking=no root@39.106.77.239 '${deploy(TAG)}'"
            }
        } else {
            echo '不需要部署[kxservd]'
        }
        echo '已结束[kxservd]部署'
    }
}

def deploy(tag) {
    def str = """
        docker pull registry-vpc.cn-beijing.aliyuncs.com/liasica/cabservd:$tag
        docker rm -f ${tag}
        mkdir -p /var/www/cabservd/${tag}/runtime
        docker run -itd --name ${tag} --restart=always \
        --network host \
        -v /var/www/cabservd/${tag}/config:/app/config \
        -v /var/www/cabservd/${tag}/runtime:/app/runtime \
        registry-vpc.cn-beijing.aliyuncs.com/liasica/cabservd:$tag
        docker image prune -f
        docker container prune -f
        docker volume prune -f
    """
    return str
}