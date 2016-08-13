package main

import (
	"fmt"
	"sort"
	"strings"
	"time"
	"unicode"
	//strconv  //json += "\\u"+strconv.FormatInt(int64(rint), 16) // json
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	n := len(s)
	if n == 0 {
		return false
	}
	if s[0] == '-' {
		s = s[1:]
	} else if s[0] == '+' {
		s = s[1:]
	}
	n = len(s)
	if n == 0 {
		return false
	}

	var isNumber = false
	i := 0
	for i < n && unicode.IsDigit(rune(s[i])) {
		i++
		isNumber = true
	}
	if i < n && s[i] == '.' {
		i++
		for i < n && unicode.IsDigit(rune(s[i])) {
			i++
			isNumber = true
		}
	}
	if isNumber && i < n && s[i] == 'e' {
		i++
		isNumber = false
		if i < n && (s[i] == '-' || s[i] == '+') {
			i++
		}
		for i < n && unicode.IsDigit(rune(s[i])) {
			i++
			isNumber = true
		}
	}
	return isNumber && i == n
}

// 对字符串数组去重
func clearRepeat(ss []string) (result []string) {
	m := make(map[string]bool)
	for _, v := range ss {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}

//利用对比相邻单词是否一样的原理来去重
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func HandingText(str string) []string {
	var ret []string
	str = strings.ToLower(str)                //转小写
	str = strings.Replace(str, ".", " ", -1)  //删除点
	str = strings.Replace(str, ",", " ", -1)  //删除逗号
	str = strings.Replace(str, "\n", " ", -1) //删除换行
	str = strings.Replace(str, "(", " ", -1)  //删除换行
	str = strings.Replace(str, ")", " ", -1)  //删除换行
	str = strings.Replace(str, ";", " ", -1)  //删除换行
	str = strings.Replace(str, "\"", " ", -1) //删除换行

	all_danci := strings.Split(str, " ") //分割成数组
	//fmt.Println(all_danci)
	sort.Strings(all_danci)
	//fmt.Println(all_danci)
	//all_danci = RemoveDuplicatesAndEmpty(all_danci)
	all_danci = clearRepeat(all_danci) //map 去重复

	number := len(all_danci) //统计单词数

	for i := 0; i < number; i++ { //循环单词
		danci := all_danci[i]
		if len(danci) > 2 && !strings.Contains(danci, "@") && !isNumber(danci) { //单词大于两位的才进来
			//fmt.Fprintf(os.Stdout, "%d %v\n", i, danci)
			ret = append(ret, danci)

		}

	}
	return ret
}
func main() {
	str := `
GitHub Terms of Service
By using the GitHub.com web site ("Service"), or any services of GitHub, Inc ("GitHub"), you are agreeing to be bound by the following terms and conditions ("Terms of Service"). IF YOU ARE ENTERING INTO THIS AGREEMENT ON BEHALF OF A COMPANY OR OTHER LEGAL ENTITY, YOU REPRESENT THAT YOU HAVE THE AUTHORITY TO BIND SUCH ENTITY, ITS AFFILIATES AND ALL USERS WHO ACCESS OUR SERVICES THROUGH YOUR ACCOUNT TO THESE TERMS AND CONDITIONS, IN WHICH CASE THE TERMS "YOU" OR "YOUR" SHALL REFER TO SUCH ENTITY, ITS AFFILIATES AND USERS ASSOCIATED WITH IT. IF YOU DO NOT HAVE SUCH AUTHORITY, OR IF YOU DO NOT AGREE WITH THESE TERMS AND CONDITIONS, YOU MUST NOT ACCEPT THIS AGREEMENT AND MAY NOT USE THE SERVICES.

Please note that if you are accessing any GitHub service in your capacity as a government entity, there are special terms that may apply to you. Please see Section G.17, below, for more details.

If GitHub makes material changes to these Terms, we will notify you by email or by posting a notice on our site before the changes are effective. Any new features that augment or enhance the current Service, including the release of new tools and resources, shall be subject to the Terms of Service. Continued use of the Service after any such changes shall constitute your consent to such changes. You can review the most current version of the Terms of Service at any time at: https://github.com/site/terms

Violation of any of the terms below will result in the termination of your Account. While GitHub prohibits such conduct and Content on the Service, you understand and agree that GitHub cannot be responsible for the Content posted on the Service and you nonetheless may be exposed to such materials. You agree to use the Service at your own risk.

A. Account Terms

You must be 13 years or older to use this Service.

You must be a human. Accounts registered by "bots" or other automated methods are not permitted.

You must provide your name, a valid email address, and any other information requested in order to complete the signup process.

Your login may only be used by one person - i.e., a single login may not be shared by multiple people - except that a machine user's actions may be directed by multiple people. You may create separate logins for as many people as your plan allows.

You are responsible for maintaining the security of your account and password. GitHub cannot and will not be liable for any loss or damage from your failure to comply with this security obligation.

You are responsible for all Content posted and activity that occurs under your account (even when Content is posted by others who have accounts under your account).

One person or legal entity may not maintain more than one free account, and one machine user account that is used exclusively for performing automated tasks.

You may not use the Service for any illegal or unauthorized purpose. You must not, in the use of the Service, violate any laws in your jurisdiction (including but not limited to copyright or trademark laws).

B. API Terms

Customers may access their GitHub account data via an API (Application Program Interface). Any use of the API, including use of the API through a third-party product that accesses GitHub, is bound by these Terms of Service plus the following specific terms:

You expressly understand and agree that GitHub shall not be liable for any direct, indirect, incidental, special, consequential or exemplary damages, including but not limited to, damages for loss of profits, goodwill, use, data or other intangible losses (even if GitHub has been advised of the possibility of such damages), resulting from your use of the API or third-party products that access data via the API.

Abuse or excessively frequent requests to GitHub via the API may result in the temporary or permanent suspension of your account's access to the API. GitHub, in its sole discretion, will determine abuse or excessive usage of the API. GitHub will make a reasonable attempt via email to warn the account owner prior to suspension.

GitHub reserves the right at any time to modify or discontinue, temporarily or permanently, your access to the API (or any part thereof) with or without notice.

C. Payment, Refunds, Upgrading and Downgrading Terms

All paid plans must enter a valid payment account. Free accounts are not required to provide payment account information.

An upgrade from the free plan to any paying plan will immediately bill you.

For monthly payment plans, the Service is billed in advance on a monthly basis and is non-refundable. There will be no refunds or credits for partial months of service, upgrade/downgrade refunds, or refunds for months unused with an open account. In order to treat everyone equally, no exceptions will be made.

When changing from a monthly billing cycle to a yearly billing cycle, GitHub will bill for a full year at the next monthly billing date.

All fees are exclusive of all taxes, levies, or duties imposed by taxing authorities, and you shall be responsible for payment of all such taxes, levies, or duties, excluding only United States (federal or state) taxes.

For any upgrade or downgrade in plan level while on a monthly billing cycle, the credit card that you provided will automatically be charged the new rate on your next billing cycle. For upgrades or downgrades while on a yearly plan, GitHub will immediately charge or refund the difference in plan cost, prorated for the remaining time in your yearly billing cycle.

Downgrading your Service may cause the loss of Content, features, or capacity of your Account. GitHub does not accept any liability for such loss.

D. Cancellation and Termination

You are solely responsible for properly canceling your account. An email or phone request to cancel your account is not considered cancellation. You can cancel your account at any time by clicking on the Account link in the global navigation bar at the top of the screen. The Account screen provides a simple no questions asked cancellation link.

All of your Content will be immediately deleted from the Service upon cancellation. This information can not be recovered once your account is cancelled.

If you cancel the Service before the end of your current paid up month, your cancellation will take effect immediately and you will not be charged again.

GitHub, in its sole discretion, has the right to suspend or terminate your account and refuse any and all current or future use of the Service, or any other GitHub service, for any reason at any time. Such termination of the Service will result in the deactivation or deletion of your Account or your access to your Account, and the forfeiture and relinquishment of all Content in your Account. GitHub reserves the right to refuse service to anyone for any reason at any time.

In the event that GitHub takes action to suspend or terminate an account, we will make a reasonable effort to provide the affected account owner with a copy of their account contents upon request, unless the account was suspended or terminated due to unlawful conduct.

E. Modifications to the Service and Prices

GitHub reserves the right at any time and from time to time to modify or discontinue, temporarily or permanently, the Service (or any part thereof) with or without notice.

Prices of all Services, including but not limited to monthly subscription plan fees to the Service, are subject to change upon 30 days notice from us. Such notice may be provided at any time by posting the changes to the GitHub Site (github.com) or the Service itself.

GitHub shall not be liable to you or to any third-party for any modification, price change, suspension or discontinuance of the Service.

F. Copyright and Content Ownership

We claim no intellectual property rights over the material you provide to the Service. Your profile and materials uploaded remain yours. However, by setting your pages to be viewed publicly, you agree to allow others to view your Content. By setting your repositories to be viewed publicly, you agree to allow others to view and fork your repositories.

GitHub does not pre-screen Content, but GitHub and its designee have the right (but not the obligation) in their sole discretion to refuse or remove any Content that is available via the Service.

You shall defend GitHub against any claim, demand, suit or proceeding made or brought against GitHub by a third-party alleging that Your Content, or Your use of the Service in violation of this Agreement, infringes or misappropriates the intellectual property rights of a third-party or violates applicable law, and shall indemnify GitHub for any damages finally awarded against, and for reasonable attorney’s fees incurred by, GitHub in connection with any such claim, demand, suit or proceeding; provided, that GitHub (a) promptly gives You written notice of the claim, demand, suit or proceeding; (b) gives You sole control of the defense and settlement of the claim, demand, suit or proceeding (provided that You may not settle any claim, demand, suit or proceeding unless the settlement unconditionally releases GitHub of all liability); and (c) provides to You all reasonable assistance, at Your expense.

The look and feel of the Service is copyright © GitHub, Inc. All rights reserved. You may not duplicate, copy, or reuse any portion of the HTML/CSS, Javascript, or visual design elements or concepts without express written permission from GitHub.

G. General Conditions

Your use of the Service is at your sole risk. The service is provided on an "as is" and "as available" basis.

Support for GitHub services is only available in English, via email.

You understand that GitHub uses third-party vendors and hosting partners to provide the necessary hardware, software, networking, storage, and related technology required to run the Service.

You must not modify, adapt or hack the Service or modify another website so as to falsely imply that it is associated with the Service, GitHub, or any other GitHub service.

You may use the GitHub Pages static hosting service solely as permitted and intended to host your organization pages, personal pages, or project pages, and for no other purpose. You may not use GitHub Pages in violation of GitHub's trademark or other rights or in violation of applicable law. GitHub reserves the right at all times to reclaim any GitHub subdomain without liability to you.

You agree not to reproduce, duplicate, copy, sell, resell or exploit any portion of the Service, use of the Service, or access to the Service without the express written permission by GitHub.

We may, but have no obligation to, remove Content and Accounts containing Content that we determine in our sole discretion are unlawful, offensive, threatening, libelous, defamatory, pornographic, obscene or otherwise objectionable or violates any party's intellectual property or these Terms of Service.

Verbal, physical, written or other abuse (including threats of abuse or retribution) of any GitHub customer, employee, member, or officer will result in immediate account termination.

You understand that the technical processing and transmission of the Service, including your Content, may be transferred unencrypted and involve (a) transmissions over various networks; and (b) changes to conform and adapt to technical requirements of connecting networks or devices.

You must not upload, post, host, or transmit unsolicited email, SMSs, or "spam" messages.

You must not transmit any worms or viruses or any code of a destructive nature.

If your bandwidth usage significantly exceeds the average bandwidth usage (as determined solely by GitHub) of other GitHub customers, we reserve the right to immediately disable your account or throttle your file hosting until you can reduce your bandwidth consumption.

GitHub does not warrant that (i) the service will meet your specific requirements, (ii) the service will be uninterrupted, timely, secure, or error-free, (iii) the results that may be obtained from the use of the service will be accurate or reliable, (iv) the quality of any products, services, information, or other material purchased or obtained by you through the service will meet your expectations, and (v) any errors in the Service will be corrected.

You expressly understand and agree that GitHub shall not be liable for any direct, indirect, incidental, special, consequential or exemplary damages, including but not limited to, damages for loss of profits, goodwill, use, data or other intangible losses (even if GitHub has been advised of the possibility of such damages), resulting from: (i) the use or the inability to use the service; (ii) the cost of procurement of substitute goods and services resulting from any goods, data, information or services purchased or obtained or messages received or transactions entered into through or from the service; (iii) unauthorized access to or alteration of your transmissions or data; (iv) statements or conduct of any third-party on the service; (v) or any other matter relating to the service.

The failure of GitHub to exercise or enforce any right or provision of the Terms of Service shall not constitute a waiver of such right or provision. The Terms of Service constitute the entire agreement between you and GitHub and govern your use of the Service, superseding any prior agreements between you and GitHub (including, but not limited to, any prior versions of the Terms of Service). You agree that these Terms of Service and Your use of the Service are governed under California law.

Questions about the Terms of Service should be sent to support@github.com.

If you are a government user or otherwise accessing or using any GitHub service in a government capacity, this Amendment to GitHub Terms of Service shall apply to you.

Map data copyright OpenStreetMap contributors and licensed under the Open Data Commons Open Database License. Map design and imagery subject to the MapBox Terms of Service.

Last Updated: February 11, 2016
    `
	t2 := time.Now()
	charlotteWeb := HandingText(str)
	fmt.Println(charlotteWeb)
	fmt.Println("单词总数：", len(charlotteWeb))
	fmt.Print("End ....\n\n")
	fmt.Println("去重用时:", time.Now().Sub(t2))

}
