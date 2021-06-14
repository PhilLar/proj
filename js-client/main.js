const Http = new XMLHttpRequest();
const url='http://127.0.0.1:3000';
Http.open("GET", url);
Http.send();

// const obj = JSON.parse(json);

// console.log(obj.count);
// // expected output: 42

// console.log(obj.result);
var obj
console.log("a")
Http.onreadystatechange = (e) => {
  obj = JSON.parse(Http.responseText);
}




ymaps.ready(function () {
    console.log(obj.FarmCharacteristics[0])


    var myMap = new ymaps.Map('map', {
            center: [59.452, 33.851],
            zoom: 9
        }, {
            searchControlProvider: 'yandex#search'
        });

        
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
 
                switch(variant){
                    case 1:

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

/////// CREATING ARRAY OF FARMS ////////
        farms = new Array()
        obj.Farms.forEach(function(entry) {
            obj.FarmCharacteristics.forEach(function(FC) {
                if (FC.FarmID == entry.ID) {
                    console.log("FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF")
                    myPlacemark = newPlaceMark(entry, FC)
                    farms.push(myPlacemark)
                } 
            })
            

            
            // console.log(entry.ID);
        });



        // farms.forEach(function(entry) {
        //     window.myObjects = ymaps.geoQuery(entry).addToMap(myMap);
        // });

////////////////// ARRAY OF REGIONS ///////////////////////

        regions = new Array()
        obj.Regions.forEach(function(entry) {
        
            KRS_farms_amount = 0
            KRS_heads = 0
            ANIMALS_farms_amount = 0
            ANIMALS_heads = 0
            SAL = 0

            obj.Farms.forEach(function(f) {
                if (f.RegionID == entry.ID) {
                    if (f.Specialization = "КРС") {
                        KRS_farms_amount++;
                    } else {
                        ANIMALS_farms_amount++;
                    }
                    KRS_heads += f.HeadsOfCows;
                    ANIMALS_heads += f.HeadsOfAnimals;
                    SAL += f.SAL;
                }
            });


            zone = new ymaps.Circle([[59.488250, 34.590600], 50000], {
                balloonContent: `<h3> Электронный паспорт района</h3>
                                    <p>Название: ${entry.Title}</p>
                                    <p>Колличество предприятий КРС: ${KRS_farms_amount}</p>
                                    <p>Колличество предприятий животных/птиц: ${ANIMALS_farms_amount}</p>
                                    <p>Поголовье коров: ${KRS_heads}</p>
                                    <p>Поголовье животных/птиц: ${ANIMALS_heads}</p>
                                    <p>Площадь земельный угодий: ${SAL}</p>`,
                hintContent: entry.Title
            }, {
                draggable: false,
                opacity: 0.001,
                fill: true,
                fillColor: "FF3333"
            });


            farmMarks = new ymaps.GeoObjectCollection();
            zone.events.add(['mouseenter', 'balloonclose', 'balloonopen', 'mouseleave'], function (e) {
                var target = e.get('target')
                type = e.get('type');
                if (type == 'mouseenter') {
                    target.options.set("opacity", 0.5);

                    
                    // myMap.addOverlay(farmMarks);
                } else if(type == 'balloonopen') {
                    var FChar
                    obj.Farms.forEach(function(farm) {
                        obj.FarmCharacteristics.forEach(function(FC) {
                            if(FC.FarmID == farm.ID) {
                                FChar = FC
                            }
                        })
                        if (farm.RegionID == entry.ID) {
                            farmMark = newPlaceMark(farm, FChar);
                            farmMarks.add(farmMark);
                            // window.myObjects = ymaps.geoQuery(farmMark).addToMap(myMap);
                        }
                    });
                    myMap.geoObjects.add(farmMarks);
                } else if (type == 'balloonclose') {
                    setTimeout(() => {  myMap.geoObjects.remove(farmMarks); }, 2000);      
                } else {
                    target.options.set("opacity", 0.001);
                }
            });

            regions.push(zone);
        });

        regions.forEach(function(entry) {
            window.myObjects = ymaps.geoQuery(entry).addToMap(myMap);
        });


        ////////////////////////////////////////////////

        // Boxitogorsk = new ymaps.Circle([[59.473576, 33.847675], 5000], {
        //     balloonContentBody: 'Балун',
        //     hintContent: 'Бокситогорск'
        // }, {
        //     draggable: false,
        //     opacity: 0.001,
        //     fill: true,
        //     fillColor: "FF3333"
        // });
        
        // Boxitogorsk.events.add(['mouseenter', 'mouseleave'], function (e) {
        //     console.log("sos")
        //     var target = e.get('target')
        //     type = e.get('type');
        //     if (type == 'mouseenter') {
        //         target.options.set("opacity", 0.5);
        //     } else {
        //         target.options.set("opacity", 0.001);
        //     }
        //     // console.log("sosi")
        //     // target.options.set("opacity", 0.85);
        //     // myCircle.options.opacity = 1;
        //     // newCircle = myCircle
        //     // newCircle.options.opacity = 1;
        //     // myCircle.options.fillColor = "33FFF6";
        //     // console.log("#########################")
        //     // console.log(myCircle)
        //     // console.log("#########################")
        //     // myMap.geoObjects.add(newCircle);
        // });

        // myMap.geoObjects.add(Boxitogorsk);
});

function newPlaceMark(farm, characteristics) {
    console.log("######################################")
    console.log(characteristics)
    console.log("######################################")
    return new ymaps.Placemark([farm.Latitude, farm.Longtitude], {
        hintContent: `<h3> Характеристика предприятия </h3>
                            <p>Название: ${farm.Title}</p>
                            <p>Масса полученного органического удобрения, т/год: ${characteristics.ManureMass}</p>
                            <p>Тип органического удобрения: ${farm.OF_type}</p>
                            <p>Масса общего азота в полученном органическом удобрении, т/год: ${characteristics.NitrogenMassInFertilizer}</p>
                            <p>Масса общего фосфора в полученном органическом удобрении, т/год т/год: ${characteristics.PhosphorMassInFertilizer}</p>
                            <p>Масса общего азота, которую готовы принять земельные угодья самого предприятия, т/год: ${characteristics.NitrogenMassForSoil}</p>
                            <p>Масса общего фосфора, которую готовы принять земельные угодья самого предприятия, т/год: ${characteristics.PhosphorMassForSoil}</p>
                            <p>Потенциал внесения органического удобрения по азоту, т/год: ${characteristics.FertilizerPotentialByNitrogen}</p>
                            <p>Потенциал внесения органического удобрения по фосфору, т/год: ${characteristics.FertilizerPotentialByPhosphor}</p>
                            <p>Необходимая дополнительная площадь земельных угодий для внесения ОУ (азот), га: ${characteristics.SquareDemandForNitrogen}</p>
                            <p>Необходимая дополнительная площадь земельных угодий для внесения ОУ (фосфор), га: ${characteristics.SquareDemandForPhosphor}</p>
                            <p>Свободная площадь земельных угодий, на которую можно внести органическое удобрение (азот) с других предприятий, га: ${characteristics.SquareFreeForNitrogen}</p>
                            <p>Свободная площадь земельных угодий, на которую можно внести органическое удобрение (фосфор) с других предприятий, га: ${characteristics.SquareFreeForPhosphor}</p>
                            <p>Необходимая вместимость навозохранилищ для переработки навоза/помета в органическое удобрение, т: ${characteristics.DemandForOFStorage}</p>`,
        // balloonContent: `"Привет, ${name}, как дела?"`
        balloonContent: `<h3> Электронный паспорт предприятия</h3>
                            <p>Название: ${farm.Title}</p>
                            <p>Специализация: ${farm.Specialization}</p>
                            <p>Поголовье животных/птиц: ${farm.HeadsOfAnimals}</p>
                            <p>Поголовье коров: ${farm.HeadsOfCows}</p>
                            <p>Площадь земельный угодий: ${farm.SAL}</p>`
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
  }