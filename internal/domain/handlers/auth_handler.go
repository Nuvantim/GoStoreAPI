package handler

import (
	"api/internal/domain/service"
	"api/pkg/utils"
	"github.com/gofiber/fiber/v2"
)

// struct login
var login struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

// struct otp
var otp struct {
	Email string `json:"email" validate:"email,required"`
}

// struct update password
var password struct {
	Otp      uint64   `json:"otp" validate:"required"`
	Password string `json:"password" validate:"required,min=8"`
}

/*
Home Handler
*/
func Home(c *fiber.Ctx) error {
	html := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1.0"><title>GoStoreAPI | HOME</title><script src="https://cdn.tailwindcss.com"></script><style>body{background-color:#111827;position:relative;overflow:hidden}body::before{content:'';position:absolute;top:0;left:0;right:0;bottom:0;z-index:-1;background:radial-gradient(circle at 20% 30%,rgba(79,70,229,0.15) 0,transparent 25%),radial-gradient(circle at 80% 70%,rgba(124,58,237,0.15) 0,transparent 25%),radial-gradient(circle at 40% 80%,rgba(109,40,217,0.15) 0,transparent 30%)}.star{position:absolute;background-color:#fff;border-radius:50%;z-index:-1;animation:twinkle var(--duration) ease-in-out infinite;opacity:0}@keyframes twinkle{0%,100%{opacity:0;transform:scale(.5)}50%{opacity:var(--opacity);transform:scale(1)}}@keyframes float{0%,100%{transform:translateY(0)}50%{transform:translateY(-20px)}}@keyframes press-glow{0%{box-shadow:0 0 0 0 rgba(109,40,217,.7)}70%{box-shadow:0 0 0 15px rgba(109,40,217,0)}100%{box-shadow:0 0 0 0 rgba(109,40,217,0)}}@keyframes shadow-pulse{0%,100%{box-shadow:0 15px 30px rgba(0,0,0,.3)}50%{box-shadow:0 25px 40px rgba(0,0,0,.4)}}@keyframes border-glow{0%,100%{border-color:rgba(199,210,254,.2)}50%{border-color:rgba(167,139,250,.4)}}.floating-box{background:rgba(17,24,39,.7);backdrop-filter:blur(12px);border:1px solid rgba(199,210,254,.2);border-radius:20px;animation:float 6s ease-in-out infinite,shadow-pulse 6s ease-in-out infinite,border-glow 8s ease-in-out infinite;transition:all .3s ease;position:relative;outline:0}.floating-box:active{animation:press-glow .8s ease-out}.floating-box:hover{animation:float 4s ease-in-out infinite,shadow-pulse 4s ease-in-out infinite,border-glow 6s ease-in-out infinite}.text-gradient{background:linear-gradient(90deg,#a78bfa 0,#c084fc 50%,#e879f9 100%);-webkit-background-clip:text;background-clip:text;color:transparent}.btn-purple-glow{background:linear-gradient(135deg,#7c3aed 0,#a855f7 50%,#c084fc 100%);box-shadow:0 4px 15px rgba(124,58,237,.4);transition:all .3s ease}.btn-purple-glow:hover{transform:translateY(-3px);box-shadow:0 8px 25px rgba(124,58,237,.6)}.bg-floating{position:absolute;border-radius:50%;background:rgba(124,58,237,.1);filter:blur(60px);z-index:-1}.floating-delay-1{animation:float 8s ease-in-out infinite 1s}.floating-delay-2{animation:float 10s ease-in-out infinite 2s}</style></head><body class="min-h-screen flex flex-col items-center justify-center p-6 text-white"><div id="stars-container"></div><div class="bg-floating w-64 h-64 top-20 left-20 floating-delay-1"></div><div class="bg-floating w-80 h-80 bottom-20 right-20 floating-delay-2"></div><div class="bg-floating w-96 h-96 top-1/3 right-1/4"></div><div class="floating-box p-8 md:p-10 max-w-2xl w-full mx-4 text-center"><h1 class="text-4xl md:text-5xl font-bold mb-6">Welcome to <span class="text-gradient">GoStoreAPI</span></h1><p class="text-lg md:text-xl mb-8 text-slate-300 leading-relaxed">GoStoreAPI is a RESTful API designed to power e-commerce platforms or store management systems. It provides structured endpoints to handle core operations such as product management, categories, transactions, user authentication, and other essential features for online stores or inventory-based applications.</p><a href="https://apistore.apidog.io" class="inline-flex items-center px-8 py-3.5 btn-purple-glow text-white font-semibold rounded-lg"><svg xmlns="http://www.w3.org/2000/svg" class="h-5 w-5 mr-2" viewBox="0 0 20 20" fill="currentColor"><path fill-rule="evenodd" d="M12.316 3.051a1 1 0 01.633 1.265l-4 12a1 1 0 11-1.898-.632l4-12a1 1 0 011.265-.633zM5.707 6.293a1 1 0 010 1.414L3.414 10l2.293 2.293a1 1 0 11-1.414 1.414l-3-3a1 1 0 010-1.414l3-3a1 1 0 011.414 0zm8.586 0a1 1 0 011.414 0l3 3a1 1 0 010 1.414l-3 3a1 1 0 11-1.414-1.414L16.586 10l-2.293-2.293a1 1 0 010-1.414z" clip-rule="evenodd"/></svg>View API Documentation</a></div><script>function createStars(){const e=document.getElementById("stars-container"),t=100,a=window.innerWidth,n=window.innerHeight;for(let o=0;o<t;o++){const t=document.createElement("div");t.classList.add("star");const l=Math.random()*2+1;t.style.width=`${l}px`,t.style.height=`${l}px`,t.style.left=`${Math.random()*a}px`,t.style.top=`${Math.random()*n}px`,t.style.setProperty("--duration",`${Math.random()*3+2}s`),t.style.setProperty("--opacity",Math.random()*.7+.3),t.style.animationDelay=`${Math.random()*5}s`,e.appendChild(t)}}const e=document.querySelector(".floating-box");e.addEventListener("mousedown",function(){this.style.animation="press-glow .8s ease-out"}),e.addEventListener("animationend",function(){this.style.animation="float 6s ease-in-out infinite, shadow-pulse 6s ease-in-out infinite, border-glow 8s ease-in-out infinite"}),window.addEventListener("load",createStars),window.addEventListener("resize",function(){document.getElementById("stars-container").innerHTML="",createStars()});</script></body></html>`
	return c.Type("html").SendString(html)
}

/*
LOGIN HANDLER
*/
func Login(c *fiber.Ctx) error {
	// Bind data
	if err := c.BodyParser(&login); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid Body Request"})
	}

	// validate data
	if err := utils.Validator(login); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}
	// Panggil service untuk login
	accessToken, refreshToken, err := service.Login(login.Email, login.Password)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{"error": err.Error()})
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken,
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
		Path:     "/",
	})

	// Kirim access token dan refresh token
	return c.JSON(fiber.Map{
		"message":      "Login Success!",
		"access_token": accessToken,
	})
}

/*
LOGOUT HANDLER
*/
func Logout(c *fiber.Ctx) error {
	// Clear the access token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "access_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   -1, // Delete cookie
	})

	// Clear the refresh token cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    "",
		HTTPOnly: true,
		Secure:   true, // Set to true in production
		SameSite: "Strict",
		Path:     "/",
		MaxAge:   -1, // Delete cookie
	})

	return c.JSON(fiber.Map{
		"message": "Logout successful",
	})
}

/*
Send OTP HANDLER
*/
func SendOTP(c *fiber.Ctx) error {
	// bind body
	if err := c.BodyParser(&otp); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid Body Request"})
	}
	send, err := service.SendOTP(otp.Email)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error,
		})
	}

	return c.Status(200).JSON(fiber.Map{"message": send})

}

func UpdatePassword(c *fiber.Ctx) error {
	// bind body
	if err := c.BodyParser(&password); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid Body Request"})
	}

	// validate data
	if err := utils.Validator(password); err != nil {
		return c.Status(422).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	update, err := service.UpdatePassword(password.Otp, password.Password)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{
			"messsage": err.Error(),
		})
	}
	service.DeleteOTP(password.Otp)
	return c.Status(200).JSON(fiber.Map{
		"message": update,
	})
}
