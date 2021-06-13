ymaps.ready(function () {
    var myMap = new ymaps.Map('map', {
            center: [59.32, 29.59],
            zoom: 9
        }, {
            searchControlProvider: 'yandex#search'
        }),

        // Создаём макет содержимого.
        MyIconContentLayout = ymaps.templateLayoutFactory.createClass(
            '<div style="color: #FFFFFF; font-weight: bold;">$[properties.iconContent]</div>'
        ),

        myPlacemark = new ymaps.Placemark([59.97, 30.31], {
            hintContent: 'Собственный значок метки',
            balloonContent: 'Это красивая метка'
        }, {
            // Опции.
            // Необходимо указать данный тип макета.
            iconLayout: 'default#image',
            // Своё изображение иконки метки.
            iconImageHref: 'img/map-marker.png',
            // Размеры метки.
            iconImageSize: [30, 42],
            // Смещение левого верхнего угла иконки относительно
            // её "ножки" (точки привязки).
            iconImageOffset: [-5, -38],
            quantity: "krs"
        });

        console.log("here0");
        function checkState () {
            // создаем переменные
            var shownObjects,
                quantity=$('#city').val(),
                // city=$('#birds').val(),
                // filter_c=new ymaps.GeoQueryResult(),
                filter_q=new ymaps.GeoQueryResult()
            
            // проверяем с какими данными мы вообще работаем и в зависимости от этого убераем ненужные метки
                var variant=0;
    
                if(quantity!=null){
                    variant+=1;
                }
                // if(city!=null){
                //     variant+=10;
                // }

                console.log("here1");
                switch(variant){
                    case 1:
                        console.log("here");
                        filter_q = myObjects.search('options.quantity="'+quantity+'"').add(filter_q);
                        shownObjects=filter_q.addToMap(myMap);
                        break;
                    // case 10:
                    //     filter_c=myObjects.search('options.city="'+city+'"').add(filter_c);
                    //     shownObjects=filter_c.addToMap(myMap);
                    //     break;
                    // case 11:
                    //     filter_q = myObjects.search('options.quantity="'+quantity+'"').add(filter_q);
                    //     filter_c=myObjects.search('options.city="'+city+'"').add(filter_c);
                    //     shownObjects=filter_c.intersect(filter_q).addToMap(myMap);
                    //     break;
                }
            
            // оставляем на карте только найденые метки
            myObjects.remove(shownObjects).removeFromMap(myMap);
        }
        
        // следим за изменениями выпадающих списков
        // $('#city').change(checkState);
        $('#city').change(checkState);




        $('#quantity').change(checkState);

        // myMap.geoObjects.add(myPlacemark)



        myPlacemark = new ymaps.Placemark([59.97, 30.31], {
            hintContent: 'Собственный значок метки',
            balloonContent: 'Это красивая метка'
        }, {
            // Опции.
            // Необходимо указать данный тип макета.
            iconLayout: 'default#image',
            // Своё изображение иконки метки.
            iconImageHref: 'img/map-marker.png',
            // Размеры метки.
            iconImageSize: [30, 42],
            // Смещение левого верхнего угла иконки относительно
            // её "ножки" (точки привязки).
            iconImageOffset: [-5, -38],
            quantity: "krs"
        });

        window.myObjects = ymaps.geoQuery(myPlacemark).addToMap(myMap);
});