package model

import (
	"time"
)

/** mysql schema
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

/** sqlite schema
CREATE TABLE `word` (
  `id` integer primary key AUTOINCREMENT,
  `origin` varchar(30) NOT NULL DEFAULT '',
  `target` varchar(250) NOT NULL DEFAULT '',
  `hits` integer NOT NULL DEFAULT '1',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
)

create unique index uni_origin on word (origin);

create trigger word_Update
before update on word
for each row
begin
     update word set updated_at = datetime('now','localtime') where id = old.id;
end;
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
