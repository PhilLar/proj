const Http = new XMLHttpRequest();
const url='http://127.0.0.1:3000';
Http.open("GET", url);
Http.send();

// const obj = JSON.parse(json);

// console.log(obj.count);
// // expected output: 42

// console.log(obj.result);

console.log("a")
Http.onreadystatechange = (e) => {
  obj = JSON.parse(Http.responseText);
  console.log(Http.responseText)
  console.log("f")
  console.log(obj[0].Title)
}



ymaps.ready(function () {
    var myMap = new ymaps.Map('map', {
            center: [59.452, 33.851],
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


        farms = new Array()
        obj.Farms.forEach(function(entry) {

            myPlacemark = new ymaps.Placemark([entry.Latitude, entry.Longtitude], {
                hintContent: entry.Title,
                // balloonContent: `"Привет, ${name}, как дела?"`
                balloonContent: `<h3> Электронный паспорт предприятия</h3>
                                    <p>Название: ${entry.Title}</p>
                                    <p>Специализация: ${entry.Specialization}</p>
                                    <p>Поголовье животных/птицы: ${entry.HeadsOfAnimals}</p>
                                    <p>Поголовье коров: ${entry.HeadsofCows}</p>
                                    <p>Площадь земельный угодий: ${entry.SAL}</p>`
                // balloonContent: `<h3 class="red">Hi, everyone!</h1>Электронный паспорт предприятия\ndfgdgd\nsdx`
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

            farms.push(myPlacemark)
            // console.log(entry.ID);
        });



        farms.forEach(function(entry) {
            window.myObjects = ymaps.geoQuery(entry).addToMap(myMap);
        });






        Boxitogorsk = new ymaps.Circle([[59.473576, 33.847675], 5000], {
            balloonContentBody: 'Балун',
            hintContent: 'Хинт'
        }, {
            draggable: false,
            opacity: 0.001,
            fill: true,
            fillColor: "FF3333"
        });
        
        Boxitogorsk.events.add(['mouseenter', 'mouseleave'], function (e) {
            console.log("sos")
            var target = e.get('target')
            type = e.get('type');
            if (type == 'mouseenter') {
                target.options.set("opacity", 0.5);
            } else {
                target.options.set("opacity", 0.001);
            }
            // console.log("sosi")
            // target.options.set("opacity", 0.85);
            // myCircle.options.opacity = 1;
            // newCircle = myCircle
            // newCircle.options.opacity = 1;
            // myCircle.options.fillColor = "33FFF6";
            // console.log("#########################")
            // console.log(myCircle)
            // console.log("#########################")
            // myMap.geoObjects.add(newCircle);
        });

        myMap.geoObjects.add(Boxitogorsk);
});