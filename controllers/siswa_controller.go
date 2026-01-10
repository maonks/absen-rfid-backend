package webcontroller

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/maonks/absen-rfid-backend/models"
	"gorm.io/gorm"
)

func SiswaPage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var siswa []models.Siswa
		db.Find(&siswa)

		return c.Render("pages/siswa_page", fiber.Map{
			"Siswa": siswa,
		}, "layouts/main")
	}
}

// // GET /siswa/create
func CreateSiswa(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var kartuKosong []models.Kartu

		db.Where("siswa_id IS NULL").Find(&kartuKosong)

		return c.Render("components/tambah_siswa", fiber.Map{
			"KartuKosong": kartuKosong,
		})
	}
}

// POST /siswa/store
func StoreSiswa(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		tgl, _ := time.Parse("2006-01-02", c.FormValue("tanggal_lahir"))

		siswa := models.Siswa{
			NIS:          c.FormValue("nis"),
			Nama:         c.FormValue("nama"),
			JenisKelamin: c.FormValue("jenis_kelamin"),
			TempatLahir:  c.FormValue("tempat_lahir"),
			TanggalLahir: tgl,
			Kelas:        c.FormValue("kelas"),
			Jurusan:      c.FormValue("jurusan"),
			Alamat:       c.FormValue("alamat"),
			NamaWali:     c.FormValue("nama_wali"),
			NoHP:         c.FormValue("no_hp"),
			Status:       c.FormValue("status"),
		}

		if err := db.Create(&siswa).Error; err != nil {
			return c.Status(400).SendString("Gagal menyimpan siswa")
		}

		// 🔥 jika kartu dipilih → update kartu
		kartuID := c.FormValue("kartu_id")
		if kartuID != "" {
			db.Model(&models.Kartu{}).
				Where("id = ?", kartuID).
				Update("siswa_id", siswa.ID)
		}
		// Tutup modal & reload halaman siswa
		return c.SendString(`
		<script>
			closeModal();
			location.reload();
		</script>
	`)
	}
}

// GET /siswa/:id/edit
func EditSiswa(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {
		id := c.Params("id")

		var siswa models.Siswa
		if err := db.First(&siswa, id).Error; err != nil {
			return c.Status(404).SendString("Siswa tidak ditemukan")
		}

		var kartuKosong []models.Kartu
		db.Where("siswa_id IS NULL").Find(&kartuKosong)

		return c.Render("components/edit_siswa", fiber.Map{
			"Siswa":       siswa,
			"KartuKosong": kartuKosong,
		})
	}
}

// POST /siswa/:id/update
func UpdateSiswa(db *gorm.DB) fiber.Handler {

	return func(c *fiber.Ctx) error {

		id := c.Params("id")

		tgl, _ := time.Parse("2006-01-02", c.FormValue("tanggal_lahir"))

		if err := db.Model(&models.Siswa{}).
			Where("id = ?", id).
			Updates(models.Siswa{
				NIS:          c.FormValue("nis"),
				Nama:         c.FormValue("nama"),
				JenisKelamin: c.FormValue("jenis_kelamin"),
				TempatLahir:  c.FormValue("tempat_lahir"),
				TanggalLahir: tgl,
				Kelas:        c.FormValue("kelas"),
				Jurusan:      c.FormValue("jurusan"),
				Alamat:       c.FormValue("alamat"),
				NamaWali:     c.FormValue("nama_wali"),
				NoHP:         c.FormValue("no_hp"),
				Status:       c.FormValue("status"),
			}).Error; err != nil {
			return c.Status(400).SendString("Gagal update siswa")
		}

		// 🔥 jika kartu dipilih → update kartu
		kartuID := c.FormValue("kartu_id")
		if kartuID != "" {
			db.Model(&models.Kartu{}).
				Where("id = ?", kartuID).
				Update("siswa_id", id)
		}

		return c.SendString(`
		<script>
			closeModal();
			location.reload();
		</script>
	`)
	}
}
