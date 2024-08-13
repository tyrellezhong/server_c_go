begin_tag=$(hostname)
begin_tag="hostname: ${begin_tag} ip: $(hostname -I)"
svrname="dsagentsvr"
echo "----------------- $begin_tag $svrname begin--------------------------"
echo "pwd=$(pwd)"
cd log/${svrname}*
ls | grep "${svrname}.ops"
tail "${svrname}.ops"
echo "                       "
echo "tencentclb config"
cd ~
cd tencentclb
cat tencentclb*.xml
echo "----------------- $begin_tag $svrname end--------------------------"