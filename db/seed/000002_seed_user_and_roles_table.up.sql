-- Insert User Roles 
INSERT INTO user_roles (id, name) VALUES
(1, 'Super Admin'),
(2, 'Admin Regional'),
(3, 'Operator Sekolah'),
(4, 'Guru');

-- Insert users
-- default password: password1234
INSERT INTO users (id, name, fullname, email, password, user_role_id, office_id, is_super_admin)
VALUES (7296478839281065984, 'Felix', 'Felix Seran', 'felix@aksara.com', '$2y$10$agn8hHbQEc9dlNhDAb.X3OFuwkdS.0oaT19/FHXs1CQtYu7WbCmge', 1, 7296845621728681984, 1), 
  (7296479361367076864, 'Admin Kab. Belu', 'Mateus Asato', 'mateus@aksara.com', '$2y$10$WN4Nr51s8SwXNFcFxBwxd.Xm68Jpdrnsk7jWAmZPrhmttUdqpNFu2', 2, 7296848305215021056, 0);


INSERT INTO users (id, name, fullname, email, password, user_role_id, office_id, school_id)  
VALUES (7296479858459213824, 'operatorSdnAsulun', 'Operator Sdn Asulun', 'asulun_sdn@aksara.com', '$2y$10$/mVyRdSlecfY6n4koBbHNe9D4Sh.EY.xxBeDPM4/QXVg6QV9IXNPO', 3, 7296848305215021056, 7296849298463956992),
  (7296481359994880000, 'guruAndri', 'Andri Raosa', 'andri_guru_asulun@aksara.com', '$2y$10$CiDkhXGEQmK4H9xgPpn1h.SLLF3y9hlLCG6eVFGqD6ciM21u3skDq', 4, 7296848305215021056, 7296849298463956992);


