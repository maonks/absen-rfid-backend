package controllers

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type HadirRow struct {
	ID     uint
	Uid    string
	Nama   string
	Masuk  string
	Pulang *string
	Status string
}

type TanpaRow struct {
	ID   uint
	Uid  string
	Nama string
}

func RiwayatPage(db *gorm.DB) fiber.Handler {
	return func(c *fiber.Ctx) error {

		var hadir []HadirRow
		var tanpa []TanpaRow

		// ================= SISWA HADIR =================
		db.Raw(`
			WITH daily AS (
			  SELECT
			    k.id AS id,
			    k.uid AS uid,
			    k.nama AS nama,
			    MIN(a.waktu) AS masuk,
			    MAX(
			      CASE
			        WHEN a.waktu::time >= '16:00:00'
			        THEN a.waktu
			      END
			    ) AS pulang
			  FROM absens a
			  JOIN kartus k ON k.uid = a.uid
			  WHERE DATE(a.waktu) = CURRENT_DATE
			  GROUP BY k.id, k.uid, k.nama
			)
			SELECT
			  id,
			  uid,
			  nama,
			  to_char(masuk, 'HH24:MI:SS') AS masuk,
			  CASE
			    WHEN pulang IS NOT NULL
			      THEN to_char(pulang, 'HH24:MI:SS')
			    ELSE NULL
			  END AS pulang,
			  CASE
			    WHEN pulang IS NULL THEN 'MASUK'
			    ELSE 'PULANG'
			  END AS status
			FROM daily
			ORDER BY nama
		`).Scan(&hadir)

		// ================= TANPA KETERANGAN =================
		db.Raw(`
			SELECT
			  k.id,
			  k.uid,
			  k.nama
			FROM kartus k
			LEFT JOIN absens a
			  ON a.uid = k.uid
			  AND DATE(a.waktu) = CURRENT_DATE
			WHERE a.uid IS NULL
			ORDER BY k.nama
		`).Scan(&tanpa)

		return c.Render("fragments/riwayat_page", fiber.Map{
			"Hadir":           hadir,
			"TanpaKeterangan": tanpa,
		}, "layouts/main")
	}
}
