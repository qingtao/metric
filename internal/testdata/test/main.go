package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"metric/euclidean"
	"runtime"
	"sort"
	"sync"
	"time"

	"gonum.org/v1/gonum/spatial/vptree"
)

type co [128]float64

func (a co) Distance(b vptree.Comparable) float64 {
	bi := b.(co)
	return euclidean.Distance(a[:], bi[:])
}

var num = flag.Int("num", 100000, "测试的数据数量")

func main() {
	flag.Parse()
	l := *num
	fmt.Println("num: ", l)
	f := make([]vptree.Comparable, l)

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < l; i++ {
		var fi co
		for j := 0; j < 128; j++ {
			fi[j] = rand.Float64()
		}
		sort.Float64s(fi[:])
		f[i] = vptree.Comparable(fi)
	}
	// b, err := json.MarshalIndent(f, "", " ")
	// if err != nil {
	// 	log.Fatalln(err)
	// 	return
	// }
	// if err = ioutil.WriteFile("./data.json", b, 0644); err != nil {
	// 	log.Fatalln(err)
	// }

	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM)
	// for sig := range ch {
	// 	switch sig {
	// 	case syscall.SIGTERM, syscall.SIGHUP:
	// 		return
	// 	default:
	// 		fmt.Println(sig.String())
	// 	}
	// }
	sort.Float64s(p[:])
	cur := time.Now()

	var dis = make([]float64, 0, l)
	for i := 0; i < len(f); i++ {
		dis = append(dis, (f[i].Distance(p)))
	}
	sort.Float64s(dis)
	since := time.Since(cur)
	fmt.Println("direct:", since.String())
	for i := 0; i < 5; i++ {
		fmt.Println(dis[i])
	}
	t, err := vptree.New(f, 5, nil)
	if err != nil {
		log.Fatalln(err)
	}
	dis = dis[:0]
	cur = time.Now()
	t.Do(func(c vptree.Comparable, i int) (done bool) {
		if n := p.Distance(c); n <= 0.162 {
			dis = append(dis, n)
		}
		return
	})
	since = time.Since(cur)
	fmt.Println("vptree:", since.String())
	sort.Float64s(dis)
	for i := 0; i < 5 && i < len(dis); i++ {
		fmt.Println(dis[i])
	}

	d := directMultiCPU(f, p)
	for i := 0; i < 5 && i < len(d); i++ {
		fmt.Println(d[i])
	}

	d = vptreeMultiCPU(f, p)
	for i := 0; i < 5 && i < len(d); i++ {
		fmt.Println(d[i])
	}
}

func directMultiCPU(f []vptree.Comparable, p co) []float64 {
	n := runtime.NumCPU()
	// n := 16
	step := len(f) / n
	wg := new(sync.WaitGroup)
	c := make(chan float64, 1024)
	dis := make([]float64, 0, len(f))
	cur := time.Now()
	go func() {
		for d := range c {
			dis = append(dis, d)
		}
	}()
	for i := 0; i < n; i++ {
		var a []vptree.Comparable
		if i == n-1 {
			a = f[i*step:]
		} else {
			a = f[i*step : (i+1)*step]
		}
		wg.Add(1)
		go func(a []vptree.Comparable, p co, c chan float64) {
			defer wg.Done()
			for i := 0; i < len(a); i++ {
				c <- a[i].Distance(p)
			}
		}(a, p, c)
	}
	wg.Wait()
	close(c)
	since := time.Since(cur)
	fmt.Println("direct with multi cpu", since.String())
	sort.Float64s(dis)
	return dis
}

func vptreeMultiCPU(f []vptree.Comparable, p co) []float64 {
	n := runtime.NumCPU()
	// n := 16
	step := len(f) / n
	wg := new(sync.WaitGroup)
	c := make(chan float64, 1024)
	ts := make([]*vptree.Tree, n)
	for i := 0; i < n; i++ {
		var a []vptree.Comparable
		if i == n-1 {
			a = f[i*step:]
		} else {
			a = f[i*step : (i+1)*step]
		}
		t, err := vptree.New(a, 5, nil)
		if err != nil {
			log.Fatalln(err)
		}
		ts[i] = t
	}
	cur := time.Now()
	dis := make([]float64, 0, len(f))

	go func() {
		for d := range c {
			dis = append(dis, d)
		}
	}()
	for i := 0; i < len(ts); i++ {
		wg.Add(1)
		go func(t *vptree.Tree, p co, c chan float64) {
			defer wg.Done()
			t.Do(func(o vptree.Comparable, i int) (done bool) {
				if n := p.Distance(o); n <= 0.162 {
					c <- n
				}
				return
			})
		}(ts[i], p, c)
	}
	wg.Wait()
	close(c)
	since := time.Since(cur)
	fmt.Println("vptree with multi cpu", since.String())
	sort.Float64s(dis)
	return dis
}

var p = co{
	0.9931382952348637,
	0.6003425197404785,
	0.08347187296969714,
	0.9314594674164455,
	0.425238047188385,
	0.3673998977516114,
	0.7088356207756185,
	0.6229611382570983,
	0.9036103935944466,
	0.899142633034247,
	0.7196621752393024,
	0.9694447059210717,
	0.367028568453427,
	0.2311690861305386,
	0.8553152550833156,
	0.2120765978050337,
	0.3815479016552494,
	0.23155406105642323,
	0.13493769893983246,
	0.4303559101676513,
	0.6401770999179125,
	0.5148870172265898,
	0.1593000139171002,
	0.5819343040401845,
	0.3635405124663262,
	0.14149064303226322,
	0.6000413081890533,
	0.6657320261524065,
	0.38520662879924916,
	0.9259265443049278,
	0.9885481060349492,
	0.36859116536026854,
	0.7965874186816414,
	0.5460249927664561,
	0.9473282423281777,
	0.48518611609392276,
	0.8911954339307926,
	0.18519119080713067,
	0.052460567033812404,
	0.9930174667329947,
	0.6978895023056644,
	0.8143648839177207,
	0.8094758629091335,
	0.43009800308435736,
	0.10139778748406075,
	0.9826379930534134,
	0.20353464788914835,
	0.483974319305199,
	0.9650041800079963,
	0.30633787713690297,
	0.8489197611454841,
	0.6563152217351358,
	0.033387677791278984,
	0.6370254499404344,
	0.9068683639590468,
	0.07448322041853789,
	0.1998008942401347,
	0.26453338600138987,
	0.5226323102855943,
	0.35570057160185414,
	0.14493703642442163,
	0.39453462356530355,
	0.9306460242949772,
	0.011308203000479692,
	0.8850764481984833,
	0.17798144161524077,
	0.24727509121654756,
	0.8396764586890924,
	0.8998627158058397,
	0.4796860101417794,
	0.49913296568259646,
	0.514864196483014,
	0.17577327435542844,
	0.669504209012038,
	0.47514798387571555,
	0.5191288593075571,
	0.8053342828356678,
	0.11961727984031582,
	0.4835939089200615,
	0.12959914355192867,
	0.8493918861149314,
	0.47693741225295366,
	0.9923298699371838,
	0.6231077047159229,
	0.3027198716236747,
	0.5404888781200355,
	0.035991754687082,
	0.14849840893747918,
	0.27329837582220523,
	0.6885872854623447,
	0.5823334643414908,
	0.7374879565974767,
	0.3621354769144847,
	0.652231186257185,
	0.9771262824154366,
	0.31452216370334857,
	0.9590172879643964,
	0.4605762696919147,
	0.9344313928632478,
	0.270775289192632,
	0.5575604837810793,
	0.5749505561371765,
	0.776789003625236,
	0.3827374982395465,
	0.9256439456828894,
	0.4163732534054975,
	0.20032481411512396,
	0.1781008324097572,
	0.6825556714653556,
	0.7893844342760966,
	0.9857295659780203,
	0.983594592059567,
	0.40681404864259474,
	0.5316151118440192,
	0.7490556932174781,
	0.2489896033313938,
	0.27694564704147456,
	0.8804016204344702,
	0.5617293393445838,
	0.59720262726195,
	0.4466639181369261,
	0.8651854828880543,
	0.4504428725246631,
	0.19025284066902107,
	0.3029861309977493,
	0.35197633523175526,
	0.07260966770542955,
	0.485994954665295,
}
