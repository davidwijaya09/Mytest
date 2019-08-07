package main
import (
	"net/http"
	"encoding/json"
)

type Data struct {
    Id string `json:"_id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Forward_link string `json:"forward_link"`
	Preview_url string `json:"preview_url"`
	Is_active bool `json:"is_active"`
	Target_device int `json:"target_device"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
	V int `json:"__v"`
	ID string `json:"id"`
	Content_ID int `json:"content_id"`
}


func main() {
 
	var musics[3] Data

	musics[0] = Data{
		Id: "5d43a5739a4ffb00231a7fb1",
		Title: "Silver Noah-Love Pledge",
		Description: "Silver Noah-Love Pledge",
		Forward_link: "http://sushiroll.co.id/smartfren/play/578",
		Preview_url: "http://sushiroll.co.id/assets/_preview/vid/___578___/c_LovePledge.jpg",
		Is_active: true,
		Target_device: 0,
		CreatedAt: "2019-08-02T02:52:35.896Z",
        UpdatedAt: "2019-08-02T02:52:35.913Z",
        V: 0,
        ID: "5d43a5739a4ffb00231a7fb1",
	}

	musics[1] = Data{
		Id: "5d43a5919a4ffb00231a7fb2",
		Title: "Top 10 Quotes Anime",
		Description: "Top 10 Quotes Anime",
		Forward_link: "http://sushiroll.co.id/smartfren/play/569",
		Preview_url: "http://sushiroll.co.id/assets/_preview/vid/___569___/Sushiroll_Jeanice.Quotes_.jpg",
		Is_active: true,
		Target_device: 0,
		CreatedAt: "2019-08-02T02:53:05.957Z",
        UpdatedAt: "2019-08-02T02:53:05.962Z",
        V: 0,
        ID: "5d43a5919a4ffb00231a7fb2",
	}

	musics[2] = Data{
		Id: "5d43a5bc9a4ffb00231a7fb3",
		Title: "Top 10 Karakters Anime",
		Description: "Top 10 Karakters Anime",
		Forward_link: "http://sushiroll.co.id/smartfren/play/568",
		Preview_url: "http://sushiroll.co.id/assets/_preview/vid/___568___/Sushiroll_Jeanice.Character_.jpg",
		Is_active: true,
		Target_device: 0,
		CreatedAt: "2019-08-02T02:53:48.277Z",
        UpdatedAt: "2019-08-02T02:53:48.285Z",
        V: 0,
        ID: "5d43a5bc9a4ffb00231a7fb3",
	}

	var assorteds[3] Data

	assorteds[0] = Data{
		Id: "5d43a5009a4ffb00231a7fae",
        Is_active: true,
        Preview_url: "https://m.smartfren.com/assets/img/news/news-339.jpg",
        Description: "Warga Batam, Nikmati Kuota Internet Smartfren 4G 60GB Cuma Rp. 60.000,-",
        Forward_link: "https://m.smartfren.com/id/news-detail/warga-batam-nikmati-kuota-internet-smartfren-4g-60gb-cuma-rp-60000--/",
        Title: "Warga Batam, Nikmati Kuota Internet Smartfren 4G 60GB Cuma Rp. 60.000,-",
        Target_device: 0,
        CreatedAt: "2019-08-02T02:50:40.395Z",
        UpdatedAt: "2019-08-02T02:50:40.420Z",
        V: 0,
        ID: "5d43a5009a4ffb00231a7fae",
	}

	assorteds[1] = Data{
		Id: "5d43a52d9a4ffb00231a7faf",
        Title: "SUPER 4G UNLIMITED",
        Forward_link: "https://m.smartfren.com/id/super-4g-unlimited",
        Preview_url: "https://m.smartfren.com/assets/img/prabayar/neo-180918/KontenWebsite1-1.jpg",
        Description: "SUPER 4G UNLIMITED",
        Is_active: true,
        Target_device: 0,
        CreatedAt: "2019-08-02T02:51:25.748Z",
        UpdatedAt: "2019-08-02T02:51:25.752Z",
        V: 0,
        ID: "5d43a52d9a4ffb00231a7faf",
	}

	assorteds[2] = Data{
		Id: "5d43a54c9a4ffb00231a7fb0",
        Title: "SUPER 4G KUOTA",
        Preview_url: "https://m.smartfren.com/assets/img/promo/super-4G-kuota/neo110419/Rev-Kartu-Perdana-Super-4G-Kuota-copy_01.png",
        Forward_link: "https://m.smartfren.com/id/super-4g-kuota",
        Description: "SUPER 4G KUOTA",
        Is_active: true,
        Target_device: 0,
        CreatedAt: "2019-08-02T02:51:56.911Z",
        UpdatedAt: "2019-08-02T02:51:56.916Z",
        V: 0,
        ID: "5d43a54c9a4ffb00231a7fb0",
	}

	var games[1] Data

	games[0] = Data{
		Id: "5d43a7cc9a4ffb00231a7fb7",
        Title: "Mobile Legends: Bang Bang",
        Preview_url: "https://lh3.googleusercontent.com/8O4uzew9XkwVQDInFhgMFuoYYw46nfaIxsT0tWeY4vk2LUcTGPvR4A_kECG0WhyJDBpl=s360-rw",
        Forward_link: "https://play.google.com/store/apps/details?id=com.mobile.legends",
        Description: "Mobile Legends: Bang Bang",
        Is_active: true,
		Target_device: 0,
		Content_ID: 0,
        CreatedAt: "2019-08-02T03:02:36.616Z",
        UpdatedAt: "2019-08-02T03:02:36.632Z",
        V: 0,
        ID: "5d43a7cc9a4ffb00231a7fb7",	
	}

	var movies[3] Data

	movies[0] = Data{
		Id: "5d43a66f9a4ffb00231a7fb4",
        Title: "Mission Impossible 2",
        Preview_url: "https://sinarmas-vod.s3-ap-southeast-1.amazonaws.com//poster//movies//1//Missionimpossible2_300x400.jpg",
        Forward_link: "http://dev.genflix.co.id//smartfren//movies//mission-impossible-2",
        Description: "Seorang agen rahasia dikirim ke Sydney, untuk menemukan dan menghancurkan virus penyakit, hasil rekayasa secara genetis yang disebut \\\"Chimera\\\"\"",
        Is_active: true,
		Target_device: 0,
        CreatedAt: "2019-08-02T02:56:47.658Z",
        UpdatedAt: "2019-08-02T02:56:47.677Z",
        V: 0,
        ID: "5d43a66f9a4ffb00231a7fb4",	
	}

	movies[1] = Data{
		Id: "5d43a6b59a4ffb00231a7fb5",
        Title: "Walk The Line",
        Preview_url: "https://sinarmas-vod.s3-ap-southeast-1.amazonaws.com//poster//movies//3//cu_WalkTheLine300R1.jpg",
        Forward_link: "http://dev.genflix.co.id//smartfren//movies//walk-the-line",
        Description: "Sebuah riwayat hidup legenda musik country, Johnny Cash.  Dimulai dari peternakan kapas Arkansas sampai ketenarannya dengan Sun Records di Memphis",
        Is_active: true,
		Target_device: 0,
        CreatedAt: "2019-08-02T02:57:57.854Z",
        UpdatedAt: "2019-08-02T02:57:57.860Z",
        V: 0,
        ID: "5d43a6b59a4ffb00231a7fb5",	
	}

	movies[2] = Data{
		Id: "5d43a6ec9a4ffb00231a7fb6",
        Title: "Miss Congeniality",
        Preview_url: "https://sinarmas-vod.s3-ap-southeast-1.amazonaws.com//poster//movies//4//MissCongeniality_300x400.jpg",
        Forward_link: "http://dev.genflix.co.id//smartfren//movies//miss-congeniality",
        Description: "Miss Amerika Serikat menjadi target pembunuh berantai, FBI memutuskan untuk menempatkan agen rahasia perempuan sebagai salah satu peserta dalam kontes kecantikan itu.",
        Is_active: true,
		Target_device: 0,
        CreatedAt: "2019-08-02T02:58:52.357Z",
        UpdatedAt: "2019-08-02T02:58:52.362Z",
        V: 0,
        ID: "5d43a6ec9a4ffb00231a7fb6",
	}

	http.HandleFunc("/Musics", func(w http.ResponseWriter, r *http.Request) {
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(musics)
	})

	
	http.HandleFunc("/Assorteds", func(w http.ResponseWriter, r *http.Request) {
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(assorteds)
	})
 
	http.HandleFunc("/Games", func(w http.ResponseWriter, r *http.Request) {
		
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(games)
	})

	http.HandleFunc("/Movies", func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(movies)
	})

	http.ListenAndServe(":16161", nil)
}

