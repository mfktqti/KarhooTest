package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"test/common"
	"time"
)

//token token
var (
	token        = &common.Token{}
	host  string = "https://rest.sandbox.karhoo.com"
)

func main() {

	/*
		文档网站：https://developer.karhoo.com

		0、获取Token
		1、刷新Token
		2. 发送询价请求(quotes)
		3. 后去询价列表(quotes)
		4. 下订单(booking)
		5. 获取订单详情
		6. 取消订单
	*/
	//账号密码
	//1、获取Token
	if err := GetToken(); err != nil {
		return
	}

	//2、刷新Token
	go RefreshToken()

	//3、询价
	res, err := RequestQuotes()
	if err != nil {
		return
	}
	//4、接收询价结果
	res, err = ReceiveQuotes(res.Id)
	if err != nil {
		return
	}

	//随机取一条报价下单
	quotesLen := len(res.Quotes)
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(quotesLen)

	qid := res.Quotes[index].Id
	fmt.Printf("随便取%s \n", res.Quotes[index].Fleet.Name)
	resQuote, err := VerifyQuotePriceCheck(qid)
	if err != nil {
		return
	}
	//下单
	resBookings, err := Bookings(resQuote.Id)
	if err != nil {
		return
	}

	bookId := resBookings.Id
	if err = GetBookingDetails(bookId); err != nil {
		return
	}

	CancelBookings(bookId)
}

//CancelBookings 取消订单
func CancelBookings(bookId string) {
	fmt.Println("取消订单-------------")
	urls := fmt.Sprintf("%s/v1/bookings/%s/cancel", host, bookId)
	data := make(map[string]interface{}, 1)
	data["reason"] = "ETA_TOO_LONG"
	_, err, _ := common.RequestPost(urls, data, token.Access_token)
	if err != nil {
		fmt.Println("取消订单出错", err)
		return
	}
	fmt.Println("取消订单成功")

}

//GetBookingDetails 获取订单详情成功
func GetBookingDetails(bookId string) error {
	fmt.Println("获取订单详情----------------")
	urls := fmt.Sprintf("%s/v1/bookings/%s", host, bookId)
	content, err, _ := common.RequestGet(urls, token.Access_token)
	if err != nil {
		fmt.Println("获取订单详细出错:", err)
		return err
	}
	res := common.Resbookings{}
	json.Unmarshal([]byte(content), &res)
	fmt.Printf("获取订单详情：%#v \n", res)
	fmt.Println("获取订单详情成功-------------")
	return nil
}

//Bookings 下单
func Bookings(quote_id string) (res common.Resbookings, err error) {
	fmt.Println("下单-------------")
	urls := fmt.Sprintf("%s/v1/bookings/", host)
	reqBookings := common.ReqBookings{}
	reqBookings.QuoteId = quote_id
	reqBookings.PassengerDetails = make([]common.PassengerDetail, 0, 1)
	reqBookings.PassengerDetails = append(reqBookings.PassengerDetails,
		common.PassengerDetail{LastName: "json", PhoneNumber: "+448000000000"})
	dataStr, _ := json.Marshal(reqBookings)
	content, err, _ := common.RequestPostStr(urls, string(dataStr), token.Access_token)
	if err != nil {
		fmt.Println("下订单出错：", err)
		return
	}
	res = common.Resbookings{}
	json.Unmarshal([]byte(content), &res)
	fmt.Printf("下单成功,单号:%s \n", res.Id)
	return
}

//VerifyQuotePriceCheck
func VerifyQuotePriceCheck(qid string) (res common.Quote, err error) {
	fmt.Println("验证有效价-------------")
	urls := fmt.Sprintf("%s/v2/quotes/verify/%s", host, qid)

	content, err, _ := common.RequestGet(urls, token.Access_token)

	if err != nil {
		fmt.Println("Verify Quote Price Check failed ,err:", err)
		return
	}
	res = common.Quote{}
	json.Unmarshal([]byte(content), &res)
	//fmt.Printf("验证结果：\n%s\n", content)
	fmt.Println("验证成功-------------")
	return
}

//ReceiveQuotes 接收询价结果
func ReceiveQuotes(id string) (res common.ResQuotes, err error) {
	fmt.Println("接收询价结果：-----------------")
	urls := fmt.Sprintf("%s/v2/quotes/%s", host, id)

	//fmt.Println(urls)
	content, err, _ := common.RequestGet(urls, token.Access_token)
	if err != nil {
		fmt.Println("Receive quotes failed,err:", err)
		return
	}
	//fmt.Println("接收询价结果 ：\n", content)
	res = common.ResQuotes{}
	json.Unmarshal([]byte(content), &res)

	fmt.Println("收到询价成功-------------")
	//fmt.Printf("%#v \n", res)
	return
}

//RequestQuotes 发送询价请求
func RequestQuotes() (res common.ResQuotes, err error) {
	fmt.Println("发送询价请求：-----------------")
	urls := fmt.Sprintf("%s/v2/quotes/", host)

	data := make(map[string]interface{}, 1)
	data["origin"] = common.LongitudeAndLatitude{Longitude: "0.0785452", Latitude: "51.5004564"}
	data["destination"] = common.LongitudeAndLatitude{Longitude: "0.0776452", Latitude: "51.5054564"}
	data["local_time_of_pickup"] = "2021-06-05T20:28"
	fmt.Println("token:", token)
	content, err, _ := common.RequestPost(urls, data, token.Access_token)
	if err != nil {
		fmt.Println("RequestQuotes failed,err:", err)
		return
	}
	res = common.ResQuotes{}
	json.Unmarshal([]byte(content), &res)
	fmt.Println("发送询价请求成功：-----------------")
	//fmt.Printf("%#v \n", res)
	//fmt.Printf("结果：%#v\n", res)
	return
}

//RefreshToken 刷新Token
func RefreshToken() {
	fmt.Println("刷新Token-----------------")
	time.Sleep(3500 * time.Second)
	urls := fmt.Sprintf("%s/v1/auth/refresh", host)

	dataRefresh := make(map[string]interface{}, 1)
	dataRefresh["refresh_token"] = token.Refresh_token

	content, err, _ := common.RequestPost(urls, dataRefresh, "")
	if err != nil {
		fmt.Println("刷新Token失败:", err)
	} else {
		fmt.Println("刷新Token成功.")
		refreshToken := common.Token{}
		err = json.Unmarshal([]byte(content), &refreshToken)
		if err != nil {
			fmt.Println("反序列化Token失败,err:", err)
		}
		token.Access_token = refreshToken.Access_token
	}
}

//GetToken 获取Token
func GetToken() error {
	fmt.Println("获取Token-----------------")
	data := make(map[string]interface{}, 2)
	data["username"] = "chuoxian.yang@external.karhoo.com"
	data["password"] = "12345678KarhooYcx"

	urls := fmt.Sprintf("%s/v1/auth/token", host)
	content, err, _ := common.RequestPost(urls, data, "")
	if err != nil {
		fmt.Println("get token failed,err:", err)
		return err
	}
	token = &common.Token{}

	err = json.Unmarshal([]byte(content), token)
	if err != nil {
		fmt.Println("反序列化Token失败,err:", err)
		return err
	}
	fmt.Println("获取Token成功")
	return nil
}
