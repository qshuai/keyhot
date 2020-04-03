package model

import (
	"time"
)

/** schema
CREATE TABLE `word` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `origin` varchar(30) NOT NULL DEFAULT '',
  `target` varchar(250) NOT NULL DEFAULT '',
  `hits` int(11) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uni_origin` (`origin`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4
*/

// Word ...
type Word struct {
	ID        int       `json:"id"`
	Origin    string    `json:"origin"`
	Target    string    `json:"target"`
	Hits      int       `json:"hits"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
