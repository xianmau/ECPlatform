$(function(){
	var v_content=$(".slide_content");
	var v_width=v_content.width();
	var v_show=$(".slide_content_list");
	var v_list=v_show.find("ul");
	var v_count=v_show.find("ul li").length;
	var index=1;
	var time=function(){
		v_show.animate({scrollLeft:'-=1144px'},"slow",function(){
		v_list.append(v_list.children().first().remove());
        v_show.scrollLeft(0);
		});
	};
	var setTime = setInterval(time,3000);
	v_content.on("mouseover",function(){
		clearInterval(setTime);
	});
	v_content.on("mouseleave",function(){
		setInterval(setTime);
	});
	//next photo
	$("span.next").click(function(){
		if(!v_show.is(":animated")){
			if(index==v_count){
				v_show.animate({left:'0px'},"slow");
				index=1;
			}else{
				v_show.animate({left:'-='+v_width},"slow");
				index++;
			}
		}
	});
	//prev photo
	$("span.prev").click(function(){
	    if(!v_show.is(":animated")){
	    	if(index==1){
	    		v_show.animate({left:'-='+v_width*(v_count-1)},"slow");
	    		index=v_count;
	    	}else{
	    		v_show.animate({left:'+='+v_width},"slow");
	    		index--;
	    	}
	    }
	});
});