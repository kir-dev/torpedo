Torpedo
=======

Torpedó játék kivonuláshoz.

Részlet az [ötletelős gistből](https://gist.github.com/tmichel/6972145):

>Klasszikus torpedó játék. Azzal bolondítjuk meg, hogy több játékos játszik
>egyszerre a szokásos 1v1 helyett. A játékba akármikor be lehet szállni, úgy
>mint egy igazi csatába. Egyszerre csak egy csata van.
>
>Van egy központi kivetítő, amin követhető a csata menete. A kivetítőn csak a már
>publikus információk láthatóak. Ilyenek például a lövések helyei és a már
>elsüllyedt hajók. Az egyes játékosokat külön színnel jelöljük. Valamint esetleg
>némi statisztika is lehet a kivetítőn:
>
>* az adott csatáról: kinek mennyi hajója van még, mennyit talált el
>* összesített: high score lista
>
>A játék kiegyensúlyozása érdekében a játékosok nem helyezhetik el a hajóikat,
>hanem egy viszonylag véletlen elhelyezést kapnak. Sőt a később csatlakozó
>játékosok csak egy csökkentett hajókészletet kapnak, hogy ne legyen akkora
>előnyük. A hajókhoz egy pontszámot rendelünk, ami a méretüktől függ. Minél
>nagyobb annál könnyebb eltalálni, de annál több körig él. Ebből a két tényezőből
>kell egy pontszámot alkotni és az újonnan csatlakozó játékosok a már bent lévő
>játékosok állapotának megfelelő leosztást kap.
>
>A pálya viszonylag nagy lesz, de azért limitált. Esetleg lehet a játékosok
>számát maximálni, hogy még értelmes maradjon a játék és a vaktában lövöldözés ne
>legyen annyira kifizetődő.

Fejlesztői környezet
--------------------

Legyen nálad Go telepítve ([link](http://golang.org/doc/install#install)). Után
töltsd le a forráskódot `go get` segítségével.

    $ go get github.com/kirdev/torpedo

Vagy állítsd be a `$GOPATH` változót és klónozd a repót, ahogy mindig is
szoktad. Utána futass `go build`-et.

Szerkeszőtnek én a Sublime Text 2/3 + GoSublime kombót ajánlom.
