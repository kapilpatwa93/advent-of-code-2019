package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

type node struct {
	isTracked bool
	children  []string
	parent    string
}


func main() {
	start := time.Now()
	input := [] string { "LZ2)Q48","91M)VLK","H7V)XCK","7LD)KQN","H9V)MDB","4MW)N1D","8SG)53Z","Q52)TL7","HV9)CD7","XVZ)GY7","PC6)1VL","NSS)QGJ","HMH)Q3D","YKG)WQ4","V1D)273","H52)3PR","VQK)FK8","WTB)ST9","1LB)MCQ","WMQ)YD4","1B7)5PR","TTY)VCZ","BN5)HNM","CMF)FF9","MCP)DTT","376)4VM","28Z)N2F","J95)H9V","JLK)C2V","NMG)FJH","FPK)1HT","J15)XPB","NFB)QWV","MFB)RBQ","3L3)KK7","LGL)48N","2CN)T62","1DN)HRX","5S8)YPD","1CK)F7Z","HJ6)YYS","3RV)X8F","8F5)8ZP","H5P)V8P","DTT)WLQ","ZB6)8LJ","NF9)MS5","KGR)DNP","85M)YS4","GWG)YLQ","KKW)J55","DRL)6R5","SQ2)M4V","K3D)3LB","BKC)5NP","KR2)L86","HLB)GTR","3C1)WNS","BPL)5XZ","8KZ)986","4V2)9FG","9TF)61T","KDN)MW9","RBR)RCK","G5V)Y8W","S4V)LRV","6Y2)5QZ","XCH)RZ2","273)PMW","F2H)L95","KFM)FFJ","7V3)79J","QGJ)QJP","ZYN)7R1","6BG)696","XVJ)X6Q","MHS)HCJ","TWY)61D","BC5)TS7","TD3)ZRH","M7N)R1N","7FY)Q69","WGN)H52","5VS)PN7","XW1)F75","ZBD)WTB","2M9)FBJ","KLC)F3K","D4S)CJ6","MLG)95F","BMG)582","133)B4C","TD3)47K","V5D)KBM","G8G)4DX","F6H)DPK","9BT)LC1","QX9)QKS","DNC)WDJ","X3V)GVT","DJM)HVV","TS7)P5M","7FG)MVP","XJV)QG3","CFR)1S5","X4Q)9KX","133)5NS","X9C)X97","HS8)Q3H","4LB)KD5","KBL)148","13G)Z1W","GPZ)B4H","72G)TQB","1HD)MJM","4YC)BY8","7YT)Q7L","MTC)44D","MDG)JF9","PHN)W77","VSJ)SLH","RGZ)GKH","NDJ)14R","S7P)H4W","LVM)L9V","8X4)5M4","G4T)H1W","X78)FZX","KQN)5MZ","3LB)WM1","ZGJ)ZT1","B4C)Z38","K1S)ZGJ","3WW)W13","BSV)WB6","N7N)BM1","YW2)QX2","CQQ)C7M","J9T)CW8","FMG)LG6","K73)B4L","DNF)91H","4K7)NRC","PWT)4G3","RRK)5S8","KXV)8WC","R65)S2P","L2J)455","GQV)YYM","LG6)ZKN","B5F)96J","VZG)DRL","J25)F26","FJ2)NX2","X9T)48T","M6V)GZZ","5QZ)5C7","W6M)19X","HZH)ZJT","W77)RW4","8X2)R7F","G4X)X9T","S4M)6BH","41L)CY2","DJD)RDT","NS2)3FT","2MM)HT7","KF5)M7N","SMN)4RP","8VM)SHZ","2H1)6H3","QTK)XD3","LXK)TNB","R39)7VD","JBQ)KZC","H1W)8Q2","VBS)L7G","C29)QNR","NCZ)MWP","LRV)DVL","BQ4)TMX","J7G)7Z2","K79)J15","9SZ)CLH","T2F)L7R","X6Q)T8N","YDP)NVN","442)MLS","62F)L76","MW9)1ZK","8PS)1D7","YLQ)Z3X","4D1)2VK","VZG)9ZL","BMG)725","K8F)8J4","RF1)VTZ","MQH)YMC","TM9)ZXR","PLJ)PRN","Z99)TWY","RYF)BXB","D3F)8NR","ZLC)Q9R","1PC)JN1","NVN)N4G","Z7Y)9TF","SR4)735","BW7)ZLV","GG1)15B","58X)SVJ","V8P)93R","T19)85T","Z1R)L2J","TRB)R6G","W81)T2Z","W5S)NNQ","NB8)P63","SM1)QGS","48X)4ZY","M3C)XYD","WLW)NS9","T2T)GZB","8V6)MK6","8LJ)R39","7ZJ)HSS","Z9M)5VS","WJL)TZ7","RR3)4VW","QD9)9T2","R72)LFS","2VK)63T","74S)MX4","MNX)DFC","NRC)1PC","LRX)F6D","VTZ)PWT","VFQ)D7T","8JT)1YF","6H3)751","P1R)KDN","WFL)4PW","JKM)R72","J2R)GF8","LRD)G1Y","DSF)G97","92X)B23","XDM)NTQ","MC2)JTP","HLK)CNN","CC2)89L","TZX)4DH","MDB)9C1","6F5)DK8","MK4)5VK","D7T)Z7K","L56)1LM","KBC)MTM","B23)CHN","9P8)5WB","KFX)RBR","61D)LRD","CKC)9YH","MVV)ZBQ","2BQ)HLB","8M7)7T8","45P)7RX","STM)7VW","KW6)2H9","2Q5)9TC","CNC)4XZ","75S)378","Y9N)ZZY","YSC)W8Y","C7M)749","S56)N65","CTD)XW1","3WC)GTP","Z2D)6JS","T91)DC2","WY2)B5F","HJJ)C3X","DFJ)NKQ","DCH)KM8","2H9)87B","6QG)J3R","LTK)XJV","XL8)PD6","3X9)JPX","CDT)LTK","7QX)72R","Z22)NK2","BL2)TTG","VNR)B2X","NF4)DQ9","T8K)X6V","NKQ)3QB","DVN)CHT","4CF)GWG","R7F)1J5","2KN)PGD","G25)1R3","14K)QL2","5WB)QGL","D51)LW8","Z7C)M7X","DTQ)XD6","ND4)13G","KDX)QXF","2JV)W4C","C2V)M9M","ZRH)PHN","T6J)MN3","6C7)TGD","3C7)CVC","9T2)DC4","KM8)8DY","KJG)2XP","1PS)XPG","284)33M","B6C)MJX","7R1)CXG","35Z)SSK","R1T)HQN","G8G)J48","ZSC)XKY","PCL)VLN","3ZM)1CK","TH8)PMP","YZP)4ZD","DPN)2VL","KK7)63Q","69C)QHX","74V)BX2","L76)XFR","N6C)BMJ","P3C)S3Q","8ZP)8PH","3ZS)CZ7","KXD)BJP","96F)8NM","MLS)44K","X97)4F9","GZB)NB8","3Q4)LD6","M2F)PMD","J3R)KXD","4YY)4QT","8DJ)CW3","NB1)W5S","T56)TQG","KMT)746","G2Q)439","6SF)WP3","44D)VVZ","B2X)GJZ","WZN)JLK","H3S)6PW","RZ3)K79","44B)KVS","44K)S49","SJ1)7QY","437)P66","BMH)2CN","9NY)YN9","QJP)NMP","WMG)V5D","G25)R1T","YBM)5LX","KPR)9BD","44Q)7SH","QYZ)YOU","HDX)8M7","F93)TBD","5M4)23F","VJ2)HT3","YM3)F93","SXY)9XB","RVR)NSP","SCS)23C","83X)MX2","8HJ)KZK","8J7)45P","69X)6CR","R3D)95D","RTV)FK2","DZL)5ST","YTY)Z45","QN3)SPF","PRC)QTG","76M)M15","FMG)GYY","KDB)WZM","3WC)1QR","6RW)FRX","COM)ZGB","ZNG)9GX","K7N)HJJ","Z7C)94S","3TB)X3Q","H4W)37S","2PM)D79","FW4)CT9","PCW)JHJ","J8G)YTY","KDN)6S3","8B7)HDT","95F)T2T","S2K)ZB2","Y8N)F7H","SK8)VBK","4BG)5KS","QYJ)PQS","78Y)W81","KM9)JWX","JHT)PF7","BGV)C4J","SSK)X6H","Z5H)5Z9","9BP)TBJ","ZVN)D29","3QS)35Z","J5J)V63","4B6)LFL","RDW)73K","JKQ)XN6","ZTC)H2J","WVY)SBS","6SC)24Z","MS5)57C","YJ1)J9T","SZB)PG7","X6H)NJ5","PLD)7C9","N63)52S","CZ7)KXZ","3CT)STM","7LN)K1R","4JJ)991","VLK)6LL","25F)WMK","SQR)HBQ","1RT)7DD","ZMG)BCS","4H8)1PS","J25)WHL","F85)M1V","DZP)R7W","NTQ)NYW","P8B)RJY","TGC)KJ2","H2S)D1S","L9V)7MD","VGM)4BX","HN2)C3G","KJ2)M87","FZ3)85M","L2J)7WS","W4C)442","B61)2GH","5ZH)MYS","RP9)2XF","DK8)7V3","LMX)27N","5H7)J2N","DVN)YM3","C81)1DP","1F7)HGW","YS4)7FW","N7N)NXP","NLH)87Q","367)J8S","YGG)NB1","QX2)Q8Y","QJ7)DJD","8NR)CC2","TYM)BKR","9L6)3R7","5PT)4BM","5NC)XS1","FWD)X3V","HMH)BQJ","SQR)X7M","F5F)96F","5Y8)9K4","3SF)TZX","TSS)LN6","RJY)MZ1","MSR)YPJ","CYX)YDP","DC2)4B6","YVW)F96","9CS)X81","FY2)284","81Y)39Q","B4C)JJP","LP3)XFF","Z83)5HF","RP3)DVN","33K)DT8","4LX)4PC","GDN)C7S","NV9)2L4","9TC)NDJ","3VT)T6J","V3L)P1X","Q7L)JBQ","TQB)76M","YT5)N9C","25F)QWT","ZL3)WHY","439)M8Z","1XJ)F46","RDG)LKY","LTM)41Y","61T)QMT","65Q)G11","VQ6)KJG","7C9)H2H","HCJ)PYV","WNK)1DF","4DH)P1H","PW4)RPJ","FTK)Q6W","B36)8CS","M87)9N2","GTP)JKP","MJZ)TD3","VZX)RWZ","RBC)2PM","99P)VFQ","7LD)JHT","GKH)YF6","HCS)BLN","7T8)ZZV","N1N)LCV","RS1)ZPQ","VWZ)HJ6","SZB)6YT","DYJ)6Y2","7MD)SMW","RDS)2YY","JMQ)WZN","NT3)R3D","PNY)2XQ","NW5)ZZX","S2F)14K","96C)1ZM","KXZ)54B","1F4)HWX","RVN)NF9","9MS)MQM","682)K3D","C4J)RH9","JW8)63L","7DD)2L3","4VM)RYF","S3Q)F9T","GTR)GYJ","GKK)Z42","BC6)FMB","6PW)Z7C","ZZY)2ZG","RPJ)YP2","6SM)MBG","9ML)24X","TNB)PNY","12G)BMH","72R)2MM","884)CNZ","6R5)DT7","1JD)QKL","9KL)SAN","2P8)ZMG","T72)KF5","F5W)RTF","8TN)YBF","2XW)CKC","N7F)979","NSP)MQH","3JM)9SZ","CZD)J8C","7VV)TYB","X81)H72","VBK)3MN","YMH)TDV","N17)QTK","X6V)XVZ","GZZ)8JT","9CS)HQ1","MK6)VR6","MCQ)9MF","F96)B7G","3SL)VXJ","15N)82S","GPW)T9J","9YF)RF4","526)NCP","JKP)NB5","F7Z)B6T","41Y)KFX","XFR)4NX","N25)2L6","W13)HN2","GJZ)QYX","8Z7)CV6","SPF)QVR","1S5)L81","TWY)THQ","RXH)5ZH","H43)WVY","HNM)3B2","S5W)35V","464)VQ6","7QY)J25","75N)2KN","JBK)7HC","RBQ)8B7","378)6SM","VHD)492","R1K)DNC","D15)69X","WQJ)44B","WQ4)4K7","SPM)F98","6NL)FTQ","ZNK)WKD","BQJ)WMG","M28)7T9","H68)FTK","LZQ)SS9","ST4)GJ2","LQH)YGG","ST8)SQR","VVP)KR1","CCM)VZW","P4V)MTC","17R)85X","1VV)7ZJ","N1D)TZN","HFZ)9XD","7RX)35T","QM3)GPW","VFP)LD7","8Y6)R1K","N6S)RJ5","N2F)Y6B","M1P)41L","5DC)3GB","H12)HTF","YMC)XK1","9QG)8F5","12G)KXN","WZ1)65Q","9M8)S92","P9V)WT6","PDJ)VFV","9FG)ST4","P3D)C4B","N9C)RK9","6HN)SR5","V8P)GLQ","KVS)FWK","XT5)HLX","H44)FWD","RDH)YD1","YBF)HKW","WM1)N23","D8V)DTL","272)LQH","7BR)WJS","7FW)LRX","YP1)T56","QGS)Y9X","LRM)B1T","H4T)8TJ","MSN)D7J","QWT)LZ2","QTG)58G","FF9)GG1","8T7)TFV","7TT)Y92","N6V)LMS","RSC)LRM","J8C)DB8","LQ6)B6C","QWH)JRV","2VZ)KYG","91T)75S","H8F)55J","73K)CQQ","R3D)JFK","V2X)S6H","L7G)DZL","JXW)8N1","HRX)W5K","5WJ)HMN","WQJ)RB8","9TC)VJ2","QNR)Z99","2V1)D3F","CBJ)L4P","V2X)GZ5","1PR)7YC","NLH)MHL","ZFN)28B","WM1)8KZ","LK5)XWY","N3C)WRX","Q33)ZD4","H22)9VV","2N9)4WS","KXV)ZNG","TVH)XT5","FBJ)LJ7","3PR)44S","VRH)191","FTQ)2Q5","X76)JVH","36V)22T","NDR)TK8","Z42)WQJ","NN4)WBS","K57)BKC","227)X3J","2YB)V28","J2D)7BF","28W)MCX","148)8SG","1VL)2TQ","1RX)Z83","5SB)K91","Q6R)4KM","26R)9TP","TK4)HYC","8RS)3LD","2VL)464","JPX)9G2","72M)LC2","J95)VFJ","PRN)XCH","TQB)2WF","ZRQ)SLB","1LM)ZLL","BQ4)D15","F46)81Y","KBB)DY1","6JS)T91","JRC)Z22","414)6NL","V4L)NPN","CW3)T2F","4PZ)P4B","S67)4V2","M4V)8VT","MTT)1VV","HDT)XFQ","CHN)K2J","G13)LMX","JRV)NSS","K3Z)L56","GDN)WFL","82S)75Z","4PL)5ZG","GJJ)X7P","RNF)5DC","HFM)7SX","YLC)TF1","P7Z)M28","7Z2)BW7","RGW)WX2","7HX)P35","3MN)39T","5TW)BWN","MLJ)1KC","BJ1)YVW","3LD)YSV","BMF)9KL","X67)J2R","JTP)Z7P","6BH)YLT","PMD)HN1","B6T)56T","CVG)GST","HKR)7YT","GXZ)XHX","ZGB)32Q","W3P)P2X","P82)NVH","MTM)367","SMW)N39","DJ7)J95","XD3)PLD","YPD)FNR","7KQ)V8N","LKY)PFB","1ZM)67F","3LD)4MW","QJW)JX1","NB1)TK9","85X)WPH","Q69)31F","9FS)72G","YLT)179","9N5)K46","Y6B)WMS","5SV)ZBD","69T)QX9","Y4P)VBS","1GD)6QG","917)R9T","RSC)ZBY","FR7)MSR","735)QM3","MWP)7LN","44S)3NV","WP4)6W9","JN3)8X4","9TF)MLJ","5KS)9VR","TK8)T1Q","LJ7)RTV","YN9)WYL","MK6)YBM","32Q)ZRC","C4B)PM8","NK2)XVJ","HRX)NMG","L8K)5SV","ZRP)33D","3B2)LX4","PYV)M6T","693)3CH","FK2)RB1","SYV)P9Z","26W)CNC","WKD)WFK","6VZ)DJ2","98C)24P","CPC)7ZL","Z19)S2F","4VF)D2D","T7T)3YR","RKB)2P8","M2G)SK8","LVM)Q62","RR1)WBC","GY7)8Z7","RTV)88Y","YSV)N7F","8F9)BN5","162)Y8N","HGS)9MS","DHT)W31","YFT)15N","TPP)F6H","5C7)T9H","6PS)KQ9","RWZ)M2G","TTT)5WJ","TXK)G2Q","TL7)MS1","DRL)N62","2H9)7QX","2N9)SVM","K1R)4JF","5XZ)QQY","Q62)JFY","JMV)3Q5","CXG)63W","1D7)MXZ","MZ7)J24","2MY)94W","PZ6)DW1","6H5)MBH","H2J)GJ5","T9T)HZW","N3Y)MYL","Q48)TBZ","3SD)76V","K73)F8D","ZNK)8DS","63L)PQ6","TLM)BV2","6ZF)B3V","Y53)TNN","Q3H)H12","D29)DW6","4ZD)N6V","98K)62F","ZLV)WMQ","V1C)VRH","JY6)YP1","YBX)KW6","VFJ)Y53","GZ5)N57","CJ6)91T","5ST)JKQ","615)N63","PMP)376","ZQL)3L3","KD5)QPM","B6C)TTY","HZ8)798","P5M)CYX","XFF)F8P","6TH)8PK","5LX)DPW","YZP)3C7","SVG)C29","954)3Q4","KTS)3SF","B4L)HGS","WMS)8RS","XD6)G5V","JDX)2V7","SVG)XPT","9XB)K57","HMN)36V","NFY)4PZ","XPB)N3Y","R6K)YW2","NS9)NJM","RTW)TX6","K6N)6B8","XKY)RNF","N23)ZM2","D35)P3D","31Q)YZP","P1H)GZC","WSN)63M","XGV)V2B","8B3)6BV","W45)ZFN","ZQL)GJQ","97P)31P","VCZ)KB8","7LD)2YB","5PD)3V8","94S)MYT","Y33)9YF","8HN)RNK","TPV)KQG","4QT)4PL","BHX)HZ8","7VD)MR5","KRG)G2H","ZDP)26X","TMX)QLK","VZW)9NY","QQ6)G4X","7NM)W6M","79J)S4D","4G3)VHD","47K)1YT","DNP)ZB6","7CP)NW5","MGG)7LD","ST9)Y87","MXD)GMG","9N2)2QB","QKZ)457","XN6)KMT","14R)QC5","CV6)F4M","9KY)6SC","N4D)G25","VLN)4XD","WBC)3WC","XFX)B2F","S2K)8VM","DPH)Y2H","8WX)J9X","WMK)L9P","ZKC)CTD","F8D)GDN","L56)DXH","M15)LRW","5KS)JMV","KYG)PDJ","RFZ)HPR","BMJ)VQK","XTV)7FG","24X)M3N","QWG)S61","NBW)MC2","GYJ)Y3K","J1X)9L6","87Q)69C","NPN)6TJ","M8Z)RXH","NX2)K1S","G2K)TDD","S73)P4V","2L6)VGF","S2C)MW7","YP2)5N6","4XD)NG5","RTW)67T","ZMS)QJW","3KS)H4V","1JD)DFJ","XCK)RZ3","9VZ)XN5","WFK)WKW","MZ1)TNH","XTV)S67","HCJ)1RX","F8P)6DM","Z78)MZ7","BX2)JY6","ZDK)KPZ","SVJ)PYS","VXK)3SL","ZKN)WNK","7T9)74P","LRW)KXV","46L)92P","M6S)HL8","S4D)1PR","YF6)Z4K","8V4)8K8","PTL)R94","HT7)W3P","6SY)FYN","1HT)8B3","2NM)LG3","HVV)QYZ","W1B)3YF","GJ4)9ML","Y87)HJT","FTR)RR1","6CH)R35","YYX)XJZ","2L4)DBY","QQY)CK3","QBW)GGB","GW9)PNJ","G2L)1KK","NB5)T19","91H)G64","C13)9XF","9MF)JDZ","DFC)4CS","8PH)Y9N","BMQ)DJM","ZBY)6CH","GJ5)8FR","WDJ)VRT","QLK)3CT","Q9R)VC4","8FR)QDM","J6D)HZH","BY3)YVV","MDG)4PJ","J4L)QKZ","MVP)615","GZC)CB3","DB8)TSS","BWZ)954","RB1)WGN","FDG)764","6B8)MLG","ZT8)Q6R","NXP)N17","HLQ)HZV","BFW)ZSC","TZK)BC5","VKN)TXK","BQJ)TVH","QVR)4JZ","B6T)GKK","V63)XQL","CC2)355","1N4)B36","58G)RDH","KBC)FZ3","34D)S4V","P1X)S2K","8Y4)ZVN","2TQ)ZDP","HWX)TJM","TVZ)K6N","D3B)PQR","31P)G2K","WRX)75N","T9J)QW5","5Z9)7Z1","YST)DHT","NW4)FPK","YG5)YKG","TC9)CWX","RNC)5Q3","Q7D)DTQ","81L)SCS","HN1)Z7Y","3YR)BKN","546)4LX","CLH)D8X","NJM)HS8","DK7)6K6","1LB)4YC","RDH)9CK","HGS)1XJ","Y2H)8F9","2Q4)P1P","LD7)ND4","D7J)546","HYC)23Z","T4C)NXZ","67F)SNX","8TF)8J7","Y8J)C13","TZ7)WLW","BY9)H43","J24)BKD","X6M)Q7D","RZ2)D51","9NY)CCF","24P)JXW","QYZ)32D","VHD)693","6NS)KPG","FP9)9QG","M1V)Z8Y","W5K)XTV","37S)78Y","47P)LZW","VMZ)98K","CLQ)HTW","YHD)JRC","7SX)FTR","9RC)F8W","JFK)5DK","N6J)RRK","HTW)VVQ","SCL)B9Y","VKN)9CS","M3N)HFW","57C)ZQL","D8X)VMZ","ZZV)CBJ","G1Y)VNR","4LS)KRG","KPG)8X2","L9M)21J","T62)28W","VTH)X67","B67)RP3","95D)B67","1DX)6PS","4XR)W2S","PKK)JLB","MS1)4VF","C2K)H2S","WNW)7HX","RXT)Q31","LFS)LTM","DT8)RF1","9YZ)X5B","W94)6C7","MGW)BMG","D1S)SVQ","TFV)5KV","3CH)CCM","TZN)2XW","FZX)VKD","4VW)J6D","5MZ)PYN","W6M)XFX","6S3)KBC","XRT)DSF","1J5)3JM","492)B5J","7RZ)ZYN","D79)4W7","9XF)7RZ","JXQ)KBB","VRT)N4D","MQM)BJ1","PLD)HDX","2WF)MVV","8S3)LVM","9VR)RB4","DT8)DFL","Z38)YHD","XWY)F5W","4T3)KFM","XN5)J4L","751)162","CPC)D6Y","F75)G13","WHY)3VT","WT6)NBZ","WG8)119","HJT)3ZM","WQ4)3X9","3FT)SR4","L7R)SJ1","LN8)KLC","9C5)XDM","39T)23L","H2H)BY9","LW8)WSC","VKL)8D7","FNR)YJ1","QKL)JN3","HKW)VXK","LV9)B97","58B)V93","TTG)1VY","MW7)BY3","K3R)JDX","SVM)5NC","3T4)7NM","4PB)G18","XK1)L8K","N4G)5H7","FNR)VSJ","CHL)WY9","M9M)9VZ","SNX)BHX","5R8)H5P","NLR)HTM","9XF)Y2J","4TK)JXQ","744)P9V","RJ5)2MY","WN7)744","8K8)C1T","G18)274","P4B)HGG","T11)4TK","3CP)SMN","1QR)VZX","XFQ)C2Y","P3D)3SQ","54F)7TT","DJ2)NS2","32H)NCZ","X9W)99P","WKW)H68","DRT)W1B","LCV)D4S","BJP)NFY","94W)QR1","35V)6FB","9CK)XNR","QPM)CHL","8J6)1DN","8DY)ZX8","SHZ)W94","YSG)TGC","ZBQ)7CT","GF8)RKT","VVQ)HLQ","CZ7)PRC","MGQ)LVS","QW5)5D2","TYB)LQ6","CY2)ZKC","29K)ZLK","KR1)FXN","4JZ)5Y8","191)CTJ","Y9X)4P5","2WN)FXL","8D7)PZJ","MJM)GY1","LC1)2NM","5HF)K3R","FMB)P3C","744)7VV","76V)3CP","B4L)P8B","S61)8FL","725)1DR","H72)44Q","DT7)V2X","3SQ)265","NCP)VZG","S25)CDT","N8V)RY8","NYW)ZTC","986)X1Q","RKT)6SF","TNH)4H8","8Q2)C7N","1H6)R8R","BKN)TVZ","FJM)2N9","JRV)GT4","F3K)9KY","ZPQ)331","MBG)LK5","SR5)L6T","ZB2)CZD","FLC)RYK","Y92)QWG","R9T)ZN9","55J)28Z","1R3)SL1","6V2)9FV","N57)SR1","W8Y)HQS","1YF)Q52","ZLK)5TW","B9Y)8HN","MYS)S73","G1G)NQ9","LTK)DPH","5HS)LGL","PQS)PC3","ZPQ)RBC","PGD)NBW","M6T)3QK","D15)GPZ","YP2)P1R","KPZ)TC6","TDV)1F4","8PK)TTT","894)H7V","S2P)3SD","P1R)Q33","CS4)RXT","CVC)KX4","9FS)BQ4","746)HKR","V93)QXS","L6T)ZWJ","TF1)MNX","RZS)QKC","ZFL)BQY","54B)TYM","B4H)ZDK","HQ1)74S","V28)LZQ","R1N)DNF","5XS)P7Z","ZRQ)RNC","2GX)6Z2","4CF)2V1","X3Q)G76","6BV)TPV","WHL)G7T","N4D)9C5","JKS)437","HLX)74V","2ZG)TW8","6DM)HFM","QXF)1S1","F7H)2WN","DBY)227","JX1)48X","NZV)1LB","XPT)T4C","991)4JH","H7H)LXL","CNN)4JJ","8NM)4YY","HGW)8FB","447)4PB","WYL)HLK","5QZ)C81","72G)BQW","979)NLS","1RX)N1N","LZW)RVR","YW2)X6M","6K6)NT3","SXR)VVP","PM8)S7P","F58)WBX","FYN)C11","9SS)K8F","6QD)DYL","YM3)2H1","GJQ)TM9","CHT)92X","3SF)4CF","DPK)J63","Y3K)L34","P98)DVX","B7G)8DK","HZW)LN8","GF5)X29","91T)C2K","J7D)VD4","FFJ)69T","B1T)2XB","P9Z)JNH","QHX)1B7","ZLL)K7N","274)8T7","X7P)MXD","DVX)S25","9YH)PKK","5F4)1DX","7WS)V1D","HTF)FY2","NQJ)884","LX4)2JV","7Z1)22N","Z45)XZV","BQY)KGR","W31)FMG","ZX8)PBT","H4V)J75","BXB)3C1","PC3)QJ7","57F)NF4","5KV)TH8","268)ZNK","VFN)H3S","PYS)26R","FGM)31Q","Z3X)J7D","X1Q)NS1","TX6)Z19","ZD4)NZV","P66)7DK","GMG)B5N","99P)RDG","RNZ)MTT","TK9)55G","F98)LSQ","QKS)KTS","C3G)PNW","7QX)R1D","4F9)ZRP","BLN)9BP","PF7)SYV","BMQ)6H5","27N)PCL","JC9)F2H","GN8)TPP","265)RSC","NQ9)1JD","ZL3)T11","FP9)ZGK","5NP)34D","B97)WSN","HT3)G23","ZN9)6VZ","7YC)DPN","X7M)T9M","2PM)CYQ","J2N)6NS","KHB)Z78","JNH)BFW","7DK)1GD","15B)3VJ","R8N)GQV","L4H)F62","XNR)6F5","9G2)SQ2","D8X)DF7","NVH)Z2D","455)KR2","4W7)MDG","P1P)JBK","1KK)TRX","764)KM9","2WF)N6C","BY9)WNW","2YY)TK4","696)BSV","JJP)KHB","P87)RQ2","MJZ)L4H","J63)FP9","WX2)H4T","ZT1)8J6","CFR)ZRQ","N65)XGV","SHM)P82","9BD)HT9","66Q)ZK2","5ZH)MFB","5VK)P98","QMT)ZL3","TDD)BMF","BQW)5F4","5PR)SCL","G76)C28","93R)TLM","MR5)GXZ","CS4)W45","SPL)KDB","TY4)JMQ","19X)7FY","33D)917","21J)2B8","RYK)PC6","SL1)2M9","MXZ)RFZ","XJ9)FW4","CLR)STT","8DS)YST","44Q)6TH","FJH)FDG","DPW)JC9","J3M)RP9","DW6)R4S","Y9N)98C","GLQ)6HN","Z99)SHM","F4M)2JS","WY9)WY2","K9X)MSN","HTM)BC6","L7G)NW4","WSC)FJM","B2F)6QD","P4Z)6V2","32D)YLC","1YT)QN3","HZV)66Q","BH9)DZP","F46)ZT8","S2C)QP5","YQM)RGZ","L4C)VWZ","4PW)9CH","MN3)GJJ","P8P)J1X","95K)VKN","BKR)897","DFX)447","4MW)XL8","67T)BWZ","28B)7KQ","TBZ)MCP","Q33)682","39Q)Y33","2ZV)NQJ","74C)9N5","S92)N6S","GCB)618","5DK)RKB","JNH)PT6","SC4)6ZF","DY1)D35","QL2)6GY","57C)X2W","FK8)RS1","CD7)LP3","NMP)NJ8","74P)9PQ","WBX)81L","8DK)MH9","SS9)6RW","4PC)F5F","48N)2DF","C11)3TB","X9C)T9T","PQT)G1G","M28)7CP","JF9)R65","3CT)3WW","Q3D)VKL","SYD)DXS","63T)XTP","T9M)72M","J55)GCB","K46)ST8","NV9)RNZ","7PT)8WX","R6K)RDS","376)J5J","S49)SZB","WB6)S4M","23Z)5HS","S7P)BPL","2MR)B61","D2D)BMQ","8LM)YSC","GVT)NV9","39T)8DJ","G23)79R","MHL)WZ1","MBH)VFN","JFY)WG8","JZD)RVN","DF7)YG5","GJ2)NMD","ZB6)2F9","VTZ)58B","PMW)TC9","XS1)SPM","RZ2)5PD","B5J)SX5","C7N)H5Z","LF9)FZF","PD6)K3Z","QHS)R6K","PFB)8Y4","2JS)DFX","6K6)96C","VRT)PQT","QP5)2Q4","331)M6S","VRR)9RC","NS1)X78","JN1)5XS","T9H)YBX","D4S)7PT","56T)46L","D79)QD9","X2W)4LB","XHX)1VB","LN6)D8V","NLJ)L3V","C2Y)1HD","QG3)P87","ZZX)VRR","YYM)JMJ","HSS)3HX","WP3)ZFL","SBS)KKW","CCF)4B7","JWX)J2D","R35)2ZV","F26)SVG","LXL)V1C","KZK)97P","WJS)P8P","6W9)6SY","7CT)YFT","2XQ)CVG","WZM)54F","PZJ)C4M","8CS)RDW","9VV)G8G","6YT)LV9","N39)D3B","NNQ)4H7","9MS)T72","NMD)QWH","YVV)4XR","FXN)VGM","Z8Y)7WF","XZV)NDR","TBD)KPR","J5J)MDP","SVQ)PG4","1BX)NN4","7WF)HMH","L9P)GN8","XFR)9BT","HQN)N8V","1VY)BHP","8WC)WN7","9XD)8TN","15Z)YT5","MJM)DYJ","6LL)PW4","2F9)4T3","75Z)7P1","P5G)8V6","78L)V4L","B4J)1N4","KXN)TY4","N4G)B4J","CB3)WJL","HBQ)4LS","3VJ)2GX","G7T)MHS","QG6)Z1S","22N)R8N","J9T)S5W","9TP)5PT","XQL)N7N","6CR)TQ1","ZWJ)YSG","6TH)7BR","53Z)1H6","B3V)HV9","D6Y)F58","92N)J8G","MX4)1F7","DXH)PVC","14K)2NY","4K7)GJ4","PT6)H8F","CK3)3ZS","897)PLJ","PNJ)J7G","7VW)CLR","QR1)P4Z","FWK)HCS","N62)PZ6","LMX)SM1","31F)CPC","NLS)25F","4ZY)SW7","Y92)414","X3J)P5G","5D2)H5J","3Q5)J3M","3VT)QQ6","ZK2)ZYR","MDP)N3C","1ZK)33K","JLB)WVW","BWN)3QS","K91)C36","C81)6HX","355)4PT","1DF)DK7","4PT)YMH","BN5)272","QTK)NLJ","Z7P)M2F","XYD)G4T","JMJ)X4Q","79R)G2L","55G)DJ7","YG5)8S3","THQ)Y8J","PYV)894","4BX)TZK","RCK)SXR","TQG)MGW","V8N)74C","8N1)SF4","L81)L9M","RQ2)QHS","23L)RZS","Y8W)K9X","CW8)SCD","Q54)8LM","7HC)MGH","LFL)YQM","BM1)1BX","TNN)QG6","R1D)TRB","4B7)JYB","26X)T8K","9ZL)2BQ","XDM)M1P","798)H44","1KC)NLR","P2X)ZMS","DNF)QYJ","MYS)8Y6","2V7)97R","NXP)Z1R","KX4)95K","YYS)CLQ","LG3)NFB","BY8)4BG","G97)SPL","C36)X76","4BM)M6V","HPR)2MR","S6H)RR3","G64)CK6","KBM)DRT","CTJ)VTH","4DX)DCH","9GX)83X","PYN)CS4","CNZ)HRS","C1T)V3L","HL8)BFT","13Y)BH9","JVH)5R8","C4M)15Z","ZRC)YYX","WY3)K73","8FL)9M8","DVL)SC4","SCD)91M","29K)3ZQ","CYQ)9P8","SR1)MJZ","Y2J)JKS","76M)PCW","9CH)S2C","6QG)275","RY8)X9W","5Q3)2QR","QGL)ZLC","6GY)8L8","BJP)9FS","R6G)S56","LD6)9H9","3V8)JZD","26W)8V4","RB4)XJ9","ZGK)H22","4XZ)CMF","ZXR)FJ2","P63)26W","BFW)RTW","GT4)X9C","TW8)FGM","J5X)XVL","C28)PF9","JDZ)1RT","HTW)BL2","HV9)8PS","582)Q54","BHP)JW8","2QB)32H","7SX)12G","3X9)VFP","Q31)4D1","QKC)29K","Z1S)13Y","5ZG)2L1","SX5)9YZ","63Q)78L","F6D)Z5H","W45)47P","STT)SYD","LZQ)Y4P","RF4)FR7","ZYR)XRT","92P)268","P4V)HFZ","9FV)FLC","4JF)JKM","GYY)M3C","PNW)17R","48T)Y49","4JH)6FV","2NY)KBL","33M)5SB","XJZ)MK4","MJX)2VZ","2JV)RGW","GGB)WY3","G2H)57F","8TJ)T7T","L86)J5X","RB8)N25","2QR)KDX","NXZ)MGQ","RW4)BGV","3YF)H7H","FYN)LF9","Z1W)92N","TGD)MGG","3HX)F85","63W)6BG","7SH)NLH","3QB)CFR","9PQ)8HJ","DFL)526","V8N)3KS","WJL)QBW","52S)3RV","SF4)WP4","179)58X","JTP)LXK","KMT)Z9M","MYT)9SS","1DR)8TF","6HX)GW9","4VM)N6J","L4P)GF5","BFT)3T4","BV2)133","H5J)SXY","FZF)L4C","4PJ)PTL",}
	orbitMap := getOrbitMap(&input)
	found := getHopCount(orbitMap)
	fmt.Println(found - 2)
	elapsed := time.Since(start)
	log.Printf("took %s", elapsed)

}

func getHopCount(orbitMap *map[string]node) int {
	parentQueue := []string{}
	hopCount := 1
	parentQueue = append(parentQueue, "YOU")
	for ; len(parentQueue) != 0; {
		currNodeName := parentQueue[0]
		currNode := (*orbitMap)[currNodeName]
		if len(parentQueue) == 1 {
			parentQueue = []string{}
		} else {
			parentQueue = parentQueue[1:]
		}
		found,count := searchInChild(orbitMap, currNode, 0)
		if !found && currNode.parent != "" {
			// insert parent of current element in the parentQueue
			currNode.isTracked = true
			(*orbitMap)[currNodeName] = currNode
			parentQueue = append(parentQueue, currNode.parent)
		} else {
			hopCount += count
			break
		}
		hopCount++
	}
	return hopCount

}
// returns ifExist,hopCount(relative count from the calling node)
func searchInChild(orbitMap *map[string]node, obj node, hopCount int) (bool,int) {
	childQueue := obj.children
	found := false
	count := 0
	for _, childName := range childQueue {
		if childName == "SAN" {
			found = true
			return found, hopCount
		}
	}
	for _, childName := range childQueue {
		child := (*orbitMap)[childName]
		if !child.isTracked {
			found, count = searchInChild(orbitMap, child, hopCount+1)
			if found {
				hopCount = count
				return found,hopCount
			}
		}

	}
	return found,hopCount
}

func getOrbitMap(input *[]string) *map[string]node {
	orbitMap := map[string]node{}
	for i := range *input {
		var str = strings.Split((*input)[i], ")")
		val, exist := orbitMap[str[0]]
		if !exist {
			node := node{
				isTracked: false,
				children:  []string{str[1]},
			}
			orbitMap[str[0]] = node
		} else {
			val.children = append(val.children, str[1])
			orbitMap[str[0]] = val
		}
		childNode := orbitMap[str[1]]
		childNode.parent = str[0]
		orbitMap[str[1]] = childNode
	}
	return &orbitMap
}
