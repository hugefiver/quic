package quic

//go:generate sh -c "./mockgen_private.sh quic mock_send_conn_test.go github.com/hugefiver/quic sendConn"
//go:generate sh -c "./mockgen_private.sh quic mock_sender_test.go github.com/hugefiver/quic sender"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_internal_test.go github.com/hugefiver/quic streamI"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_stream_test.go github.com/hugefiver/quic cryptoStream"
//go:generate sh -c "./mockgen_private.sh quic mock_receive_stream_internal_test.go github.com/hugefiver/quic receiveStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_send_stream_internal_test.go github.com/hugefiver/quic sendStreamI"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_sender_test.go github.com/hugefiver/quic streamSender"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_getter_test.go github.com/hugefiver/quic streamGetter"
//go:generate sh -c "./mockgen_private.sh quic mock_crypto_data_handler_test.go github.com/hugefiver/quic cryptoDataHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_frame_source_test.go github.com/hugefiver/quic frameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_ack_frame_source_test.go github.com/hugefiver/quic ackFrameSource"
//go:generate sh -c "./mockgen_private.sh quic mock_stream_manager_test.go github.com/hugefiver/quic streamManager"
//go:generate sh -c "./mockgen_private.sh quic mock_sealing_manager_test.go github.com/hugefiver/quic sealingManager"
//go:generate sh -c "./mockgen_private.sh quic mock_unpacker_test.go github.com/hugefiver/quic unpacker"
//go:generate sh -c "./mockgen_private.sh quic mock_packer_test.go github.com/hugefiver/quic packer"
//go:generate sh -c "./mockgen_private.sh quic mock_mtu_discoverer_test.go github.com/hugefiver/quic mtuDiscoverer"
//go:generate sh -c "./mockgen_private.sh quic mock_session_runner_test.go github.com/hugefiver/quic sessionRunner"
//go:generate sh -c "./mockgen_private.sh quic mock_quic_session_test.go github.com/hugefiver/quic quicSession"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_test.go github.com/hugefiver/quic packetHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_unknown_packet_handler_test.go github.com/hugefiver/quic unknownPacketHandler"
//go:generate sh -c "./mockgen_private.sh quic mock_packet_handler_manager_test.go github.com/hugefiver/quic packetHandlerManager"
//go:generate sh -c "./mockgen_private.sh quic mock_multiplexer_test.go github.com/hugefiver/quic multiplexer"
//go:generate sh -c "./mockgen_private.sh quic mock_batch_conn_test.go github.com/hugefiver/quic batchConn"
//go:generate sh -c "mockgen -package quic -self_package github.com/hugefiver/quic -destination mock_token_store_test.go github.com/hugefiver/quic TokenStore && goimports -w mock_token_store_test.go"
//go:generate sh -c "mockgen -package quic -self_package github.com/hugefiver/quic -destination mock_packetconn_test.go net PacketConn && goimports -w mock_packetconn_test.go"
