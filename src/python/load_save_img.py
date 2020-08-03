import cv2
 
# read image as grey scale
grey_img = cv2.imread('./org.jpg', cv2.IMREAD_GRAYSCALE)
 
# save image
status = cv2.imwrite('./grey.jpg',grey_img)
 
print("Image written to file-system : ",status)