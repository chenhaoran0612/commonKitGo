package tools

import "zeus/app/kit/global"

func PageOffset(pageNum int32, pageSize int32 , count int32 ) (start int32, end int32){
	if pageSize <=0 {
		pageSize = global.DEFAULT_PAGE_SIZE
	}

	if pageNum <= 0{
		start = 0
	}else{
		start = (pageNum-1)*pageSize
		if start > count{
			start = count - pageSize
		}
	}


	end = start + pageSize
	if end > count{
		end = count
	}

	if start<0{start = 0 }

	return start , end
}
