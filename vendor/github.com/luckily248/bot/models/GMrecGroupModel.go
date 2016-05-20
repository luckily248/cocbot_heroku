//{
//  "id": "1234567890",
//  "name": "Family",
//  "type": "private",
//  "description": "Coolest Family Ever",
//  "image_url": "http://i.groupme.com/123456789",
//  "creator_user_id": "1234567890",
//  "created_at": 1302623328,
//  "updated_at": 1302623328,
//  "members": [
//    {
//      "user_id": "1234567890",
//      "nickname": "Jane",
//      "muted": false,
//      "image_url": "http://i.groupme.com/123456789"
//    }
//  ],
//  "share_url": "https://groupme.com/join_group/1234567890/SHARE_TOKEN",
//  "messages": {
//    "count": 100,
//    "last_message_id": "1234567890",
//    "last_message_created_at": 1302623328,
//    "preview": {
//      "nickname": "Jane",
//      "text": "Hello world",
//      "image_url": "http://i.groupme.com/123456789",
//      "attachments": [
//        {
//          "type": "image",
//          "url": "http://i.groupme.com/123456789"
//        },
//        {
//          "type": "image",
//          "url": "http://i.groupme.com/123456789"
//        },
//        {
//          "type": "location",
//          "lat": "40.738206",
//          "lng": "-73.993285",
//          "name": "GroupMe HQ"
//        },
//        {
//          "type": "split",
//          "token": "SPLIT_TOKEN"
//        },
//        {
//          "type": "emoji",
//          "placeholder": "☃",
//          "charmap": [
//            [
//              1,
//              42
//            ],
//            [
//              2,
//              34
//            ]
//          ]
//        }
//      ]
//    }
//  }
//}

package models

type GMrecGroupModel struct {
	Id              string
	Name            string
	Type            string
	Description     string
	Image_url       string
	Creator_user_id string
	Created_at      string
	Updated_at      string
	Members         []Member
	Share_url       string
}
type Member struct {
	User_id   string
	Nickname  string
	muted     bool
	Image_url string
}
