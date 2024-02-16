// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"

	"github.com/bitmagnet-io/bitmagnet/internal/protocol"
)

const TableNameTorrentPieces = "torrent_pieces"

// TorrentPieces mapped from table <torrent_pieces>
type TorrentPieces struct {
	InfoHash    protocol.ID `gorm:"column:info_hash;primaryKey;<-:create" json:"infoHash"`
	PieceLength int64       `gorm:"column:piece_length;not null" json:"pieceLength"`
	Pieces      []byte      `gorm:"column:pieces;not null" json:"-"`
	CreatedAt   time.Time   `gorm:"column:created_at;not null;<-:create" json:"createdAt"`
}

// TableName TorrentPieces's table name
func (*TorrentPieces) TableName() string {
	return TableNameTorrentPieces
}
