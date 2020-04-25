(function($){
	"use strict";
	$('[data-bg-image]').each(function(){
		$(this).css({ 'background-image':'url('+$(this).data('bg-image')+')' });
	});
	//border-image
	

	$('[data-width]').each(function(){
		$(this).css({ 'width': $(this).data('width') });
	});

	$('[data-height]').each(function(){
		$(this).css({ 'height': $(this).data('height')});
	});
	// header search action
	$('#header-search').on('click', function(){
		$('#overlay-search').addClass('active');
		setTimeout(function(){
			$('#overlay-search').find('input').eq(0).focus();
		}, 400);
	});
	$('#overlay-search').find('.close-search').on('click', function(){
		$('#overlay-search').removeClass('active');
	});
	// mobile menu
	$('nav.main-nav').on('click', function(){
		if( $(window).width()<997 ){
			$('nav.main-nav').addClass('show-menu');
		}
	});
	
	$('#close-menu').on('click', function(){
		$('nav.main-nav.show-menu').removeClass('show-menu');
		return false;
	});

	var _bubble_left = $('#tpl-bubble-left').html();
	var _bubble_right = $('#tpl-bubble-right').html();
	$('.bubble-line').each(function(){
		$(this).append( $('<div class="left-shape"></div>').append(_bubble_left) );
		$(this).append( $('<div class="right-shape"></div>').append(_bubble_right) );
	});


	$(document).ready(function(){
		// Gallery slideshow


		// fullwidth post slider

		$('.footer-slide').each(function(){
			var $this = $(this);
			$this.find('.swiper-container').swiper({
				slidesPerView: 6,
				nextButton: $this.find('.swiper-button-next'),
				prevButton: $this.find('.swiper-button-prev'),
				paginationClickable: true,
				paginationType: 'progress',
				breakpoints: {
			    	996: {
			    		slidesPerView: 4
			    	},
			    	720: {
			    		slidesPerView: 2
			    	}
			    }
			});
		});

		//content slide
		$('.swiper-post').each(function(){
			var $this = $(this);
			$this.find('.swiper-slide1').swiper({
				slidesPerView: 1,
				nextButton: $this.find('.swiper-button-next'),
				prevButton: $this.find('.swiper-button-prev'),
				paginationClickable: true,
				paginationType: 'progress'
			});
		});
		
		$('.s-slide').each(function(){
			var _slide = $(this);
			_slide.find('.swiper-container').swiper({
				slidesPerView: 1,
				nextButton: _slide.find('.swiper-button-next'),
				prevButton: _slide.find('.swiper-button-prev'),
				paginationClickable: true
			});
		});
		$('.content_slide').each(function(){
			var _slide = $(this);
			_slide.find('.swiper-container').swiper({
				slidesPerView: 1,
				nextButton: _slide.find('.swiper-button-next'),
				prevButton: _slide.find('.swiper-button-prev'),
				paginationClickable: true
			});
		});

		$(document).ready(function() {
			function getRandomSize(min, max) {
				return Math.round(Math.random() * (max - min) + min);
			}
			for (var i = 0; i < 2; i++) {
				var width = getRandomSize(200, 400);
			   	var height =  getRandomSize(200, 400);
				$('.photos').append('');
			}
		});




		
	});

})(jQuery);




