package main

const (
	LoginPath          = "/default.aspx"
	HomePath           = "/default.aspx"
	CourseListPath     = "/Default.aspx?page=dkmonhoc"
	RegisterCoursePath = "/ajaxpro/EduSoft.Web.UC.DangKyMonHoc,EduSoft.Web.ashx"
	SaveCoursePath     = "/ajaxpro/EduSoft.Web.UC.DangKyMonHoc,EduSoft.Web.ashx"
)

const (
	SessionIDCookieField = "ASP.NET_SessionId"
)

const (
	UserGreetingSelector = "#ctl00_Header1_Logout1_lblNguoiDung"
	CourseAlertSelector  = "#ContentPlaceHolder1_ctl00_lblThongBaoNgoaiTGDK"
)

const (
	IDInputName          = "ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$txtTaiKhoa"
	PasswordInputName    = "ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$txtMatKhau"
	LoginActionInputName = "ctl00$ContentPlaceHolder1$ctl00$ucDangNhap$btnDangNhap"
)

const (
	RegisterCourseAjaxMethod = "LuuVaoKetQuaDangKy"
	SaveCourseAjaxMethod     = "LuuDanhSachDangKy_HopLe"
)

const (
	SystemFailureMessage         = "System error 😢 "
	LoginSuccessMessage          = "Login successfully!!! 😆 "
	LoginFailureMessage          = "Login failed 😢 "
	LogoutSuccessMessage         = "Logut successfully!!! 😆 "
	LogoutFailureMessage         = "Logout failed 😢 "
	RegistrationIsOpenMessage    = "Registration is open 😆 "
	RegistrationIsNotOpenMessage = "Registration is not open 😢 "
	RegistrationSuccessMessage   = "Registered 😆 "
	RegistrationFailureMessage   = "Register failed 😢 "
	SaveSuccessMessage           = "Saved!! 😆 "
	SaveFailureMessage           = "Save failed 😢 "
)
