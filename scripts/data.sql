
-- default inserts
insert into kinopoisk.films
values (default, 'adventures', 'Стражи галактики', 'Guardians of the Galaxy', 'nlysXqG-gbQ', 0.0, 0.0, 0, 0,
        'Description', '/static/img/A1.jpg', '/static/img/A01.jpg', 'Россия', 2010, 13),
       (default, 'horrors', 'Очень страшное кино', 'Scary Movie', 'nIW4y4w502M', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A2.jpg', '/static/img/A20.jpg', 'Россия', 2015, 13),
       (default, 'war', 'Т-34', 'T-34', 'DrtRCIQVHnA', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A3.jpg', '/static/img/A30.jpg', 'Россия', 2017, 13),
       (default, 'historical', 'Принц Персии', 'Prince of Persia', 'uOCYsMJHmKk', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A4.jpg', '/static/img/A40.jpg', 'Россия', 2019, 14),
       (default, 'animated', 'Твое имя', 'Kimi no Na wa', 'tT7b5wR0IOM', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A5.jpg', '/static/img/A50.jpg', 'Россия', 2017, 11),
       (default, 'detectives', 'Убийство в восточном экспрессе', 'Murder on the Orient Express', 'pTK0hUqzolU', 0.0,
        0.0, 0, 0, 'Description', '/static/img/A6.jpg', '/static/img/60.jpg', 'Россия', 2017, 11),
       (default, 'biographical', 'Господин Никто', 'Mr. Nobody', 'aAWGDN2S-sE', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A7.jpg', '/static/img/A70.jpg', 'Россия', 2017, 18),
       (default, 'documentary', 'Он вам не Димон', 'He is not Dimon to You', 'qrwlk7_GF9g', 0.0, 0.0, 0, 0,
        'Description', '/static/img/A8.jpg', '/static/img/A80.jpg', 'Россия', 2017, 18),
       (default, 'criminal', 'Джокер', 'Joker', 'jGfiPs9zuhE', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A9.jpg', '/static/img/A90.jpg', 'Россия', 2017, 18),
       (default, 'action', 'Бригада', 'Brigada', 'l3F5Tu1AZUU', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A10.jpg', '/static/img/A100.jpg', 'Россия', 2017, 18),
       (default, 'drama', 'Титаник', 'Titanic', 'ZQ6klONCq4s', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A11.jpg', '/static/img/A110.jpg', 'Россия', 2017, 12),
       (default, 'melodrama', 'Три метра над уровнем неба', 'Three Steps Above The Sky', 'elLW9pxj4kM', 0.0, 0.0, 0, 0,
        'Description', '/static/img/A12.jpg', '/static/img/120.jpg', 'Россия', 2017, 18),
       (default, 'comedy', 'Третий лишний', 'TED', 'tJIcP7YGhw0', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/A13.jpg', '/static/img/A130.jpg', 'Россия', 2017, 18);

insert into kinopoisk.genres
values ('Приключения', 'adventures'),
       ('Ужасы', 'horros'),
       ('Военные', 'war'),
       ('Исторические', 'historical'),
       ('Анимация', 'animated'),
       ('Детективы', 'detectives'),
       ('Биографические', 'biographical'),
       ('Документальные', 'documentary'),
       ('Криминал', 'criminal'),
       ('Боевики', 'action'),
       ('Драмы', 'drama'),
       ('Мелодрамы', 'melodrama'),
       ('Комедии', 'comedy');

insert into kinopoisk.films_genres
values (1, 'animated'),
       (1, 'adventures'),
       (2, 'horros'),
       (3, 'war'),
       (4, 'historical'),
       (5, 'animated'),
       (6, 'detectives'),
       (7, 'biographical'),
       (8, 'documentary'),
       (9, 'criminal'),
       (10, 'action'),
       (11, 'drama'),
       (12, 'melodrama'),
       (13, 'comedy');

insert into kinopoisk.series
values (default, 'adventures', 'Время приключений', 'Adventure Time', '594sVuwYTKQ', 0.0, 0.0, 0, 0,
        'Description', '/static/img/1.jpg', '/static/img/01.jpg', 'Россия', 2010, 2011, 6),
       (default, 'horrors', 'Американская история ужасов', 'American Horror Story', '6Vw-xG8nLyE', 0.0, 0.0, 0, 0,
        'Description', '/static/img/2.jpg', '/static/img/20.jpg', 'Россия', 2015, 2016, 16),
       (default, 'war', 'Братья по оружию', 'Band of Brothers', 'panok5dLHM4', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/3.jpg', '/static/img/30.jpg', 'Россия', 2017, 2018, 16),
       (default, 'historical', 'Чернобыль', 'Chernobyl', 'qtY2sel76qo', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/4.jpg', '/static/img/40.jpg', 'Россия', 2019, 2020, 16),
       (default, 'animated', 'Наруто', 'Naruto', 'xTOJ5_RKdl8', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/5.jpg', '/static/img/50.jpg', 'Россия', 2017, 2018, 8),
       (default, 'detectives', 'Шерлок', 'Sherlock', 'eMM7sX4-6gc', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/6.jpg', '/static/img/60.jpg', 'Россия', 2017, 2018, 12),
       (default, 'biographical', 'Высоцкий.Четыре часа настоящей жизни', 'Vysotsky', 'wTbqwQbLmOA', 0.0, 0.0, 0, 0,
        'Description', '/static/img/7.jpg', '/static/img/70.jpg', 'Россия', 2017, 2018, 16),
       (default, 'documentary', 'Как устроена наша планета', 'How our Earth is build', 'SgApoHS6eJE', 0.0, 0.0, 0, 0,
        'Description', '/static/img/8.jpg', '/static/img/80.jpg', 'Россия', 2017, 2018, 16),
       (default, 'criminal', 'Острые козырьки', 'Peaky Blinders', '0InieEzg5kY', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/9.jpg', '/static/img/90.jpg', 'Россия', 2017, 2018, 16),
       (default, 'action', 'Флеш', 'The Flash', 'rNfU1myyZYo', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/10.jpg', '/static/img/100.jpg', 'Россия', 2017, 2018, 16),
       (default, 'drama', 'Сверхъестественное', 'Supernatural', 'la_XCx06Ric', 0.0, 0.0, 0, 0, 'Description',
        '/static/img/11.jpg', '/static/img/110.jpg', 'Россия', 2017, 2018, 18),
       (default, 'melodrama', 'Однажды в сказке', 'Once Upon A Time', 'JV4v-Lu3NUI', 0.0, 0.0, 0, 0,
        'Description', '/static/img/120.jpg', '/static/img/12.jpg', 'Россия', 2017, 2018, 18),
       (default, 'comedy', 'Теория большого взрыва', 'The Big Bang Theory', '48p90POPC5I', 0.0, 0.0, 0, 0,
        'Description', '/static/img/13.jpg', '/static/img/130.jpg', 'Россия', 2017, 2018, 18);

insert into kinopoisk.series_genres
values (1, 'animated'),
       (1, 'adventures'),
       (2, 'horros'),
       (3, 'war'),
       (4, 'historical'),
       (5, 'animated'),
       (6, 'detectives'),
       (7, 'biographical'),
       (8, 'documentary'),
       (9, 'criminal'),
       (10, 'action'),
       (11, 'drama'),
       (12, 'melodrama'),
       (13, 'comedy');

insert into kinopoisk.seasons
values (default, 1, 'season1', 1, 'link1', 'desc1', 2010, 'img1'),
       (default, 2, 'season2', 2, 'link2', 'desc2', 2010, 'img2'),
       (default, 3, 'season3', 3, 'link3', 'desc3', 2011, 'img3'),
       (default, 4, 'season21', 1, 'link21', 'desc21', 2011, 'img21');

insert into kinopoisk.episodes
values (default, 1, 'ep11', 1, 'img1'),
       (default, 1, 'ep12', 2, 'img2'),
       (default, 2, 'ep21', 1, 'img3'),
       (default, 3, 'ep31', 1, 'img3'),
       (default, 4, 'ep41', 1, 'img21');

insert into kinopoisk.persons
values (default, 'ivan ivanov', 'actor', '2020-01-01', 'Moscow');
insert into kinopoisk.persons
values (default, 'alex alexov', 'actor', '2020-01-01', 'Moscow');

insert into kinopoisk.film_actor
values (default, 1, 1);
insert into kinopoisk.film_actor
values (default, 2, 1);
insert into kinopoisk.film_actor
values (default, 3, 1);
insert into kinopoisk.film_actor
values (default, 4, 1);
insert into kinopoisk.film_actor
values (default, 5, 1);
insert into kinopoisk.film_actor
values (default, 6, 1);
insert into kinopoisk.film_actor
values (default, 7, 1);
insert into kinopoisk.film_actor
values (default, 8, 1);
insert into kinopoisk.film_actor
values (default, 9, 1);
insert into kinopoisk.film_actor
values (default, 10, 1);
insert into kinopoisk.film_actor
values (default, 11, 1);
insert into kinopoisk.film_actor
values (default, 12, 1);
insert into kinopoisk.film_actor
values (default, 13, 1);

insert into kinopoisk.film_actor
values (default, 1, 2);
insert into kinopoisk.film_actor
values (default, 2, 2);
insert into kinopoisk.film_actor
values (default, 3, 2);
insert into kinopoisk.film_actor
values (default, 4, 2);
insert into kinopoisk.film_actor
values (default, 5, 2);
insert into kinopoisk.film_actor
values (default, 6, 2);
insert into kinopoisk.film_actor
values (default, 7, 2);
insert into kinopoisk.film_actor
values (default, 8, 2);
insert into kinopoisk.film_actor
values (default, 9, 2);
insert into kinopoisk.film_actor
values (default, 10, 2);
insert into kinopoisk.film_actor
values (default, 11, 2);
insert into kinopoisk.film_actor
values (default, 12, 2);
insert into kinopoisk.film_actor
values (default, 13, 2);


insert into kinopoisk.series_actor
values (default, 1, 2);
insert into kinopoisk.series_actor
values (default, 2, 2);
insert into kinopoisk.series_actor
values (default, 3, 2);
insert into kinopoisk.series_actor
values (default, 4, 2);
insert into kinopoisk.series_actor
values (default, 5, 2);
insert into kinopoisk.series_actor
values (default, 6, 2);
insert into kinopoisk.series_actor
values (default, 7, 2);
insert into kinopoisk.series_actor
values (default, 8, 2);
insert into kinopoisk.series_actor
values (default, 9, 2);
insert into kinopoisk.series_actor
values (default, 10, 2);
insert into kinopoisk.series_actor
values (default, 11, 2);
insert into kinopoisk.series_actor
values (default, 12, 2);
insert into kinopoisk.series_actor
values (default, 13, 2);

insert into kinopoisk.series_actor
values (default, 1, 3);
insert into kinopoisk.series_actor
values (default, 2, 3);
insert into kinopoisk.series_actor
values (default, 3, 3);
insert into kinopoisk.series_actor
values (default, 4, 3);
insert into kinopoisk.series_actor
values (default, 5, 3);
insert into kinopoisk.series_actor
values (default, 6, 3);
insert into kinopoisk.series_actor
values (default, 7, 3);
insert into kinopoisk.series_actor
values (default, 8, 3);
insert into kinopoisk.series_actor
values (default, 9, 3);
insert into kinopoisk.series_actor
values (default, 10, 3);
insert into kinopoisk.series_actor
values (default, 11, 3);
insert into kinopoisk.series_actor
values (default, 12, 3);
insert into kinopoisk.series_actor
values (default, 13, 3);

update kinopoisk.persons
set image = '/static/img/person1.jpg'
where id = 1;
update kinopoisk.persons
set image = '/static/img/person2.jpg'
where id = 2;
update kinopoisk.persons
set image = '/static/img/person3.jpg'
where id = 3;

insert into kinopoisk.persons
values (default, 'Александр Паль', 'actor', '1988-12-16', 'Челябинск', '/static/img/person4.jpg');
insert into kinopoisk.persons
values (default, 'Том Харди', 'actor', '1977-09-15', 'Лондон', '/static/img/person5.jpg');
insert into kinopoisk.persons
values (default, 'Любовь Аксенова', 'actor', '1990-03-15', 'Москва', '/static/img/person6.jpg');
insert into kinopoisk.persons
values (default, 'Александра Бортич', 'actor', '1994-09-24', 'Светлогорск', '/static/img/person7.jpg');
insert into kinopoisk.persons
values (default, 'Ана де Армас', 'actor', '1988-04-30', 'Санта-Крус-дель-Норте', '/static/img/person8.jpg');
insert into kinopoisk.persons
values (default, 'Александра Даддарио', 'actor', '1986-03-16', 'Нью-Йорк', '/static/img/person9.jpg');
insert into kinopoisk.persons
values (default, 'Ева Грин', 'actor', '1980-07-06', 'Париж', '/static/img/person10.jpg');
