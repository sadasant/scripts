var net = require('net');
var TCP_DELIMITER = '|';
var packetHeaderLen = 4; // 32 bit integer --> 4

var server = net.createServer( function(c) {
    var accumulatingBuffer = new Buffer(0);
    var totalPacketLen   = -1;
    var accumulatingLen  =  0;
    var recvedThisTimeLen=  0;
    var remoteAddress = c.remoteAddress;    var address= c.address();
    var remotePort= c.remotePort;    var remoteIpPort = remoteAddress +":"+ remotePort;

    console.log('-------------------------------'+remoteAddress);
    console.log('remoteIpPort='+ remoteIpPort);

    c.on('data', function(data) {        console.log('received data length :' + data.length );
        console.log('data='+ data);

        recvedThisTimeLen = data.length;
        console.log('recvedThisTimeLen='+ recvedThisTimeLen);

        //accumulate incoming data
        var tmpBuffer = new Buffer( accumulatingLen + recvedThisTimeLen );
        accumulatingBuffer.copy(tmpBuffer);
        data.copy ( tmpBuffer, accumulatingLen  ); // offset for accumulating
        accumulatingBuffer = tmpBuffer;
        tmpBuffer = null;
        accumulatingLen += recvedThisTimeLen ;
        console.log('accumulatingBuffer = ' + accumulatingBuffer  );
        console.log('accumulatingLen    =' + accumulatingLen );

        if( recvedThisTimeLen < packetHeaderLen ) {
            console.log('need to get more data(less than header-length received) -> wait..');
            return;
        } else if( recvedThisTimeLen == packetHeaderLen ) {
            console.log('need to get more data(only header-info is available) -> wait..');
            return;
        } else {
            console.log('before-totalPacketLen=' + totalPacketLen );
            //a packet info is available..
            if( totalPacketLen < 0 ) {
                totalPacketLen = accumulatingBuffer.readUInt32BE(0) ;
                console.log('totalPacketLen=' + totalPacketLen );
            }
        }

        //while=> 
        //in case of the accumulatingBuffer has multiple 'header and message'.
        while( accumulatingLen >= totalPacketLen + packetHeaderLen ) {
            console.log( 'accumulatingBuffer= ' + accumulatingBuffer );

            var aPacketBufExceptHeader = new Buffer( totalPacketLen  ); // a whole packet is available...
            console.log( 'aPacketBufExceptHeader len= ' + aPacketBufExceptHeader.length );
            accumulatingBuffer.copy( aPacketBufExceptHeader, 0, packetHeaderLen, accumulatingBuffer.length); // 

            ////////////////////////////////////////////////////////////////////
            //process one packet data
            var stringData = aPacketBufExceptHeader.toString();
            var usage = stringData.substring(0,stringData.indexOf(TCP_DELIMITER));
            console.log("usage: " + usage);
            //call handler
            (serverFunctions [usage])(c, remoteIpPort, stringData.substring(1+stringData.indexOf(TCP_DELIMITER)));
            ////////////////////////////////////////////////////////////////////

            //rebuild buffer
            var newBufRebuild = new Buffer( accumulatingBuffer.length );
            newBufRebuild.fill();
            accumulatingBuffer.copy( newBufRebuild, 0, totalPacketLen + packetHeaderLen, accumulatingBuffer.length  );

            //init
            accumulatingLen -= (totalPacketLen +4) ;
            accumulatingBuffer = newBufRebuild;
            newBufRebuild = null;
            totalPacketLen = -1;
            console.log( 'Init: accumulatingBuffer= ' + accumulatingBuffer );
            console.log( '      accumulatingLen   = ' + accumulatingLen );

            if( accumulatingLen <= packetHeaderLen ) {
                return;
            } else {
                totalPacketLen = accumulatingBuffer.readUInt32BE(0) ;
                console.log('totalPacketLen=' + totalPacketLen );
                console.log(accumulatingBuffer.toString('utf8', 0, totalPacketLen));
            }
        }
    });

}).listen(80);
