package models

type Verse struct {
	ID              uint   `json:"id" gorm:"not null;primaryKey"`
	ChapterID       uint   `json:"cahpter_id "`
	VerseKey        string `json:"verse_key"`
	HizbNumber      uint   `json:"hizb_number"`
	RubElHizbNumber uint   `json:"rub_el_hizb_number"`
	RukuNumber      uint   `json:"ruku_number"`
}

// {
//   "id": 1,
//   "chapter_id": 1,
//   "verse_key": "1:1",
//   "hizb_number": 1,
//   "rub_el_hizb_number": 1,
//   "ruku_number": 1,
//   "manzil_number": 1,
//   "page_number": 1,
//   "juz_number": 1,
//   "text_uthmani": "بِسْمِ ٱللَّهِ ٱلرَّحْمَـٰنِ ٱلرَّحِيمِ",
//   "audio_url": "https://download.quranicaudio.com/qdc/mishari_al_afasy/murattal/1.mp3",
//   "words": [
//     {
//       "id": 1,
//       "line_number": 2,
//       "audio_url": "https://audio.qurancdn.com/wbw/001_001_001.mp3",
//       "text": "بِسْمِ",
//       "translation": "dengan nama",
//       "transliteration": "bis'mi"
//     },
//     {
//       "id": 2,
//       "line_number": 2,
//       "audio_url": "https://audio.qurancdn.com/wbw/001_001_002.mp3",
//       "text": "ٱللَّهِ",
//       "translation": "Allah",
//       "transliteration": "l-lahi"
//     },
//     {
//       "id": 3,
//       "line_number": 2,
//       "audio_url": "https://audio.qurancdn.com/wbw/001_001_003.mp3",
//       "text": "ٱلرَّحْمَـٰنِ",
//       "translation": "Maha Pengasih",
//       "transliteration": "l-raḥmāni"
//     },
//     {
//       "id": 4,
//       "line_number": 2,
//       "audio_url": "https://audio.qurancdn.com/wbw/001_001_004.mp3",
//       "text": "ٱلرَّحِيمِ",
//       "translation": "Maha Penyayang",
//       "transliteration": "l-raḥīmi"
//     },
//     {
//       "id": 5,
//       "line_number": 2,
//       "audio_url": null,
//       "text": "١",
//       "translation": "(1)",
//       "transliteration": null
//     }
//   ],
//   "translation": {
//     "text": "Dengan menyebut nama Allah Yang Maha Pemurah lagi Maha Penyayang<sup foot_note_id=\"1\">1</sup>.",
//     "resource_name": "King Fahad Quran Complex"
//   },
//   "transliteration": "Bismillāhir-raḥmānir-raḥīm(i)."
// }
