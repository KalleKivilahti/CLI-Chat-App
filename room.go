package main

import "net"

type room struct {
	name    string
	members map[net.Addr]*client
}

func (r *room) broadcast(sender *client, msg string) {
	for addr, message := range r.members {
		if sender.conn.RemoteAddr() != addr {
			message.msg(msg)
		}
	}
}
