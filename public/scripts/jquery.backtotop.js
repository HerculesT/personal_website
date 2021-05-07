jQuery("#backtotop").click(function () {
    jQuery("body,html").animate({
        scrollTop: 0
    }, 600);
});
jQuery("#contact_btn").click(function () {
    jQuery("body,html").animate({
        scrollTop: ($('#contactMe').offset().top)
    }, 900);
});


//image pop up
$(document).ready(function() {
    $('.image-link').magnificPopup({type:'image'});
  });

jQuery(window).scroll(function () {
    if (jQuery(window).scrollTop() > 150) {
        jQuery("#backtotop").addClass("visible");
    } else {
        jQuery("#backtotop").removeClass("visible");
    }
});

