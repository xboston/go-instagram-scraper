# go-instagram

# TODO
 - поиск: https://www.instagram.com/query/?q=ig_user(27105452){id,username,external_url,full_name,profile_pic_url,biography,followed_by{count},follows{count},media{count},is_private,is_verified}
 - подписки: https://www.instagram.com/query/?q=ig_user(27105452)%20{%20followed_by.first(10)%20{%20count,%20page_info%20{%20end_cursor,%20has_next_page%20},%20nodes%20{%20id,%20is_verified,%20followed_by_viewer,%20requested_by_viewer,%20full_name,%20profile_pic_url,%20username%20}%20}%20}
 - подписчики: https://www.instagram.com/query/?q=ig_user(27105452)%20{%20follows.first(10)%20{%20count,%20page_info%20{%20end_cursor,%20has_next_page%20},%20nodes%20{%20id,%20is_verified,%20followed_by_viewer,%20requested_by_viewer,%20full_name,%20profile_pic_url,%20username%20}%20}%20}
 - комментарии: https://www.instagram.com/query/?q=ig_shortcode(BG3Iz-No1IZ){comments.last(300){count,nodes{id,created_at,text,user{id,profile_pic_url,username,follows{count},followed_by{count},biography,full_name,media{count},is_private,external_url,is_verified}},page_info}}
