package message

import (
	"it-chain/service/domain"
	"fmt"
)

func PeerTableToTable(peerTable domain.PeerTable) *Message_PeerTable{

	peerTable_ := &Message_PeerTable{}
	peerTable_.PeerTable = &PeerTable{}
	peerTable_.PeerTable.OwnerID = peerTable.OwnerID
	pbPeerInfos := make([]*PeerInfo,0)

	for _, peerInfo := range peerTable.PeerMap {
		pbPeerInfo := &PeerInfo{}
		pbPeerInfo.IpAddress = peerInfo.IpAddress
		pbPeerInfo.PeerID = peerInfo.PeerID
		pbPeerInfo.HeartBeat = int32(peerInfo.HeartBeat)
		pbPeerInfo.Port = peerInfo.Port
		pbPeerInfo.PubKey = peerInfo.PubKey
		pbPeerInfos = append(pbPeerInfos, pbPeerInfo)
	}

	peerTable_.PeerTable.PeerInfos = pbPeerInfos

	return peerTable_
}

func (m *Message) FromPeerTable(peerTable domain.PeerTable){
	m.GetPeerTable().OwnerID = peerTable.OwnerID

}

func (m *Message) SetPeerInfos(peerInfos []domain.PeerInfo, ownerID string){

	fmt.Println(m.GetPeerTable())

	m.GetPeerTable().OwnerID = ownerID

	pbPeerInfos := make([]*PeerInfo,0)

	for _, peerInfo := range peerInfos {
		pbPeerInfo := &PeerInfo{}
		pbPeerInfo.PeerID = peerInfo.PeerID
		pbPeerInfo.HeartBeat = int32(peerInfo.HeartBeat)
		pbPeerInfo.Port = peerInfo.Port
		pbPeerInfo.PubKey = peerInfo.PubKey
		pbPeerInfos = append(pbPeerInfos, pbPeerInfo)
	}

	m.GetPeerTable().PeerInfos = pbPeerInfos

}

func (m *Message) GetPeerTable_() *domain.PeerTable{

	peerTable := &domain.PeerTable{}
	peerTable.OwnerID = m.GetPeerTable().OwnerID
	peerTable.PeerMap = make(map[string]*domain.PeerInfo)

	for _, peerInfo := range m.GetPeerTable().PeerInfos {
		peerInfo_ := &domain.PeerInfo{}
		peerInfo_.IpAddress = peerInfo.IpAddress
		peerInfo_.PeerID  = peerInfo.PeerID
		peerInfo_.PubKey = peerInfo.PubKey
		peerInfo_.HeartBeat = int(peerInfo.HeartBeat)
		peerInfo_.Port = peerInfo.Port

		peerTable.PeerMap[peerInfo.PeerID] = peerInfo_
	}

	return peerTable
}
