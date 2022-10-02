package logging

//go:generate sh -c "mockgen -package logging -self_package github.com/hugefiver/quic/logging -destination mock_connection_tracer_test.go github.com/hugefiver/quic/logging ConnectionTracer"
//go:generate sh -c "mockgen -package logging -self_package github.com/hugefiver/quic/logging -destination mock_tracer_test.go github.com/hugefiver/quic/logging Tracer"
