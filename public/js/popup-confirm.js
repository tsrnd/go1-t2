var totalPrice = 0
var totalQuantity = 0;
jQuery.each(data, function(key, val){
    totalPrice += val.TotalPrice
    totalQuantity += val.Quantity
    jQuery('#grand-total').text(totalPrice);
    jQuery('#subcar-total').text(totalPrice)
    jQuery('#count-cart').text(totalQuantity)
})
jQuery('body').on('click', 'a.remove-cart', function (){
    var id = jQuery(this).data('id');
    var totalPriceRemove = jQuery(this).data('total-price');
    var grandTotal = parseFloat(jQuery('#grand-total').text()) - totalPriceRemove
    var that = this;
    var totalCart = 0;
    bootbox.confirm({ 
        message: "<h4>Are you sure you want to delete?</h4>", 
        callback: function(result){
            if(result) {
                jQuery.ajax({
                    url: base_url+ '/remove-cart',
                    type: "POST",
                    data: {detailCartId: id},
                    success:function (response) {
                        if (response.statusCode == 1) {
                            jQuery(that).parent().parent().remove();
                            jQuery('#sub-cart-'+id).remove();
                            jQuery('#grand-total').text(grandTotal);
                            jQuery('#subcar-total').text(grandTotal);
                            jQuery('.quantity').each(function (keQ, valQ) {
                                totalCart+= parseInt(jQuery(this).val());
                            })
                            jQuery('#count-cart').text(totalCart);
                        } else {
                            console.log("error");
                        }
                    },
                    error:function (error) {
                        console.log(error)
                    }
                });
            }
        }
    })
    
})

// update cart
jQuery('body').on('click', 'a.update-cart', function (){
    var id = jQuery(this).data('id');
    var price = jQuery(this).data('price');
    var quantityOld = jQuery(this).data('quantity');
    var quantity = parseInt(jQuery(this).closest('tr').find('.quantity').val());
    if (quantity == "" || quantity == 0) {
        alert('Vui lòng nhập số lượng lớn hơn 0')
        jQuery(this).closest('tr').find('.quantity').val(quantityOld);
    } else {
        var that = this;
        var totalPrice = price * quantity
        var grandTotal = 0;
        var totalCartEdit = 0;
        jQuery.ajax({
            url: base_url+ '/update-cart',
            type: "POST",
            data: {detailCartId: id, quantity: quantity, totalPrice: totalPrice},
            success:function (response) {
                console.log(response);
                if (response.statusCode == 1) {
                    jQuery(that).closest('tr').find('.total-price').text(totalPrice);
                    jQuery('.total-price').each(function(){
                        grandTotal += parseFloat (jQuery(this).text());
                    });
                    jQuery('#grand-total').text(grandTotal);
                    jQuery('#quantity-'+id).text(quantity);
                    jQuery('#total-price-'+id).text(totalPrice);
                    jQuery('#subcar-total').text(grandTotal);
                    jQuery('.quantity').each(function (keQ, valQ) {
                        totalCartEdit+= parseInt(jQuery(this).val());
                    })
                    jQuery('#count-cart').text(totalCartEdit);
                } else {
                    console.log("error");
                }
            },
            error:function (error) {
                console.log(error)
            }
        });

    }
})
