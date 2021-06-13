ymaps.ready(init);
    function init() {
        var myMap = new ymaps.Map('map', {
                center: [59.91229229104291,30.3467767304687],
                zoom: 8
            });
        
        function checkState () {
            // создаем переменные
            var shownObjects,
                quantity=$('#quantity').val(),
                city=$('#city').val(),
                location=$('#location').val(),
                filter_c=new ymaps.GeoQueryResult(),
                filter_l=new ymaps.GeoQueryResult(),
                filter_q=new ymaps.GeoQueryResult();
            
            // проверяем с какими данными мы вообще работаем и в зависимости от этого убераем ненужные метки
                var variant=0;
    
                if(quantity!=null){
                    variant+=1;
                }
                if(city!=null){
                    variant+=10;
                }
                if(location!=null){
                    variant+=100;
                }
                
                switch(variant){
                    case 1:
                        filter_q = myObjects.search('options.quantity="'+quantity+'"').add(filter_q);
                        shownObjects=filter_q.addToMap(myMap);
                        break;
                    case 10:
                        filter_c=myObjects.search('options.city="'+city+'"').add(filter_c);
                        shownObjects=filter_c.addToMap(myMap);
                        break;
                    case 100:
                        filter_l = myObjects.search('options.location="'+location+'"').add(filter_l);
                        shownObjects=filter_l.addToMap(myMap);
                        break;
                    case 11:
                        filter_q = myObjects.search('options.quantity="'+quantity+'"').add(filter_q);
                        filter_c=myObjects.search('options.city="'+city+'"').add(filter_c);
                        shownObjects=filter_c.intersect(filter_q).addToMap(myMap);
                        break;
                    case 101:
                        filter_q = myObjects.search('options.quantity="'+quantity+'"').add(filter_q);
                        filter_l = myObjects.search('options.location="'+location+'"').add(filter_l);
                        shownObjects=filter_l.intersect(filter_q).addToMap(myMap);
                        break;
                    case 110:
                        filter_l = myObjects.search('options.location="'+location+'"').add(filter_l);
                        filter_c=myObjects.search('options.city="'+city+'"').add(filter_c);
                        shownObjects=filter_c.intersect(filter_l).addToMap(myMap);
                        break;
                    case 111:
                        filter_q = myObjects.search('options.quantity="'+quantity+'"').add(filter_q);
                        filter_l = myObjects.search('options.location="'+location+'"').add(filter_l);
                        filter_c=myObjects.search('options.city="'+city+'"').add(filter_c);
                        shownObjects=filter_c.intersect(filter_l).intersect(filter_q).addToMap(myMap);
                        break;
                }
            
            // оставляем на карте только найденые метки
            myObjects.remove(shownObjects).removeFromMap(myMap);
        }
        
        // следим за изменениями выпадающих списков
        $('#city').change(checkState);
        $('#location').change(checkState);
        $('#quantity').change(checkState);


        window.myObjects = ymaps.geoQuery({
            type: "FeatureCollection",
            features: [
                // здесь метки согласно ссылке ниже
            ]
        }).addToMap(myMap);
    }