INSERT INTO offices (
  id,
  code,
  name,
  province_id,
  regency_id,
  district_id,
  email,
  phone,
  logo_url,
  created_by
)
VALUES (
  1,
  'KP01',
  'Kantor Pusat',
  31,
  3171,
  317103,
  'kantor_pusat@cerdascek.com',
  '+62',
  '',
  7296478839281065984
),
(
  2,
  'KDB01',
  'Kantor Kab. Belu',
  53,
  5304,
  530412,
  'kd_belu01@cerdascek.com',
  '+62',
  '',
  7296478839281065984
);


INSERT INTO schools (
  id,
  code,
  name,
  office_id,
  province_id,
  regency_id,
  district_id,
  created_by
)
VALUES (
  1,
  'KDBLSDIA1',
  'SDI Asuulun',
  7296848305215021056,
  53,
  5304,
  530412,
  7296478839281065984
),
(
  2,
  'KDBLSMPNTB1',
  'SMPN 1 Tasifeto Barat',
  7296848305215021056,
  53,
  5304,
  530404,
  7296478839281065984
);