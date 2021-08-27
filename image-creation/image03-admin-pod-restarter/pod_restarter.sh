podlistfile=/tmp/podlist
KUBECONFIG=~/.kube/admin.kubeconfig.ksn3
logfile=/tmp/pod_restarter.log

restart_pod() {

echo "INFO - In restart_pod()" >> $logfile 2>> $logfile

}

while true
do

podcount=`kubectl get pod goweb1-1 -n default --kubeconfig=$KUBECONFIG | grep Running |grep -v grep | wc | awk '{print $1}'`
podcontainerready=`kubectl get pod goweb1-1 -n default --kubeconfig=$KUBECONFIG | grep Running |grep -v grep | awk '{print $2}'`

#echo $podcount
#echo $podcontainerready

containersready=`echo $podcontainerready | cut -d/ -f1`
containerstotal=`echo $podcontainerready | cut -d/ -f2`

if [ $podcount -eq 0 ]
then
  restart_pod
else
  if [ "${containersready}" != "${containerstotal}" ]
  then
      restart_pod
  else
      echo "INFO - Pods are fine" >> $logfile 2>> $logfile
  fi
fi

sleep 60

done
