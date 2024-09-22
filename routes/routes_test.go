package routes

import "testing"

func TestServerSetup(t *testing.T) {
	Addr := ":8080"

	server := ServerSetup(Addr)

	if server.Addr != Addr {
		t.Errorf("Server address %s does not equal the correct value %s", server.Addr, Addr)
	}

}
