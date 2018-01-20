#!/bin/bash


BIN=./bin/mpi-ms
LOGFILE=./log

# num of nodes
#NODES=8
HOSTS=30
HOSTLIST=ppc651,ppc652,ppc653,ppc654,ppc655,ppc656,ppc657,ppc658,ppc659,ppc660,ppc661,ppc662,ppc663,ppc664,ppc665,ppc666,ppc667,ppc668,ppc669,ppc670,ppc671,ppc672,ppc673,ppc674,ppc675,ppc676,ppc677,ppc678,ppc679,ppc680

# init sizes
# VOLUME == FACTOR^n * CHUNK
VOLUME=$((128 * 1024 * 1024))
# 8388608 4194304 2097152 1048576 524288 262144 131072 65536"
#CHUNKSIZES="134217728 67108864 33554432 16777216 8388608 4194304 2097152 1048576 524288 262144 131072 65536 32768 16384"
FACTOR=2


#for NODES in `seq 8 -4 2`
#for NODES in `seq 2 -1 1`
for NODES in 8
do

# reset chunk size
CHUNK=1048576

#while [ ${CHUNK} -lt ${VOLUME} ]
#do

for CACHE in 150 200 300 400 500 700 1000 2000 4000 6000 7000 8000 9000 10000 12000 16000 32000
do
    date
    echo "Berechne: [$HOSTS]-$NODES $CHUNK $VOLUME - C$CACHE"

    # logging
    LOGSTRING=`hostname | tr -d "\n" ; echo -n " - "; date | tr -d "\n"`

    mpirun -H $HOSTLIST -n $NODES $BIN -c $CHUNK -v $VOLUME -C $CACHE -m >>"data/$HOSTS-$NODES-$VOLUME-$CHUNK-$CACHE.log"

    # logging
#    echo "$LOGSTRING: $HOSTS-$NODES-$VOLUME-$CHUNK:$CACHE">>$LOGFILE

#    CHUNK=$((CHUNK*FACTOR))
# end of for chunksizes
done

# end of while chunk size
#done

# end of for NODES
done

